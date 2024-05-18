package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	"sync"
	"time"
	holiday "unicomer-test/internal/holiday/application"
	"unicomer-test/internal/holiday/application/find"
	"unicomer-test/internal/holiday/domain"
)

type Holidays struct {
	wg         sync.WaitGroup
	holidayApp *holiday.App
}

func NewEndpoint(holidayApp *holiday.App) *Holidays {
	return &Holidays{
		holidayApp: holidayApp,
	}
}
func (h *Holidays) Invoke(c *gin.Context) error {
	ctx := c.Request.Context()

	typeFilter := c.Query("type")
	startDate := c.Query("start")
	endDate := c.Query("end")
	acceptHeader := c.GetHeader("Accept")

	holidays, err := h.FilterHolidays(ctx, typeFilter, startDate, endDate)
	if err != nil {
		log.Printf("Failed to filter holidays: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter holidays"})
		return err
	}

	if acceptHeader == "application/xml" {
		c.XML(http.StatusOK, holidays)
	} else {
		c.JSON(http.StatusOK, gin.H{"holidays": holidays})
	}

	return nil
}

func (h *Holidays) FilterHolidays(ctx context.Context, typeFilter, startDate, endDate string) ([]domain.Holiday, error) {

	allHolidays, err := h.holidayApp.Queries.FindCodebase.Handle(ctx, find.HolidayFindQuery{HolidayType: "", Date: time.Time{}})
	if err != nil {
		return nil, err
	}

	filteredHolidays := make([]domain.Holiday, 0)
	for _, holiday := range allHolidays {
		if (typeFilter == "" || strings.ToLower(holiday.Type) == strings.ToLower(typeFilter)) &&
			(startDate == "" || holiday.Date.After(parseDate(startDate))) &&
			(endDate == "" || holiday.Date.Before(parseDate(endDate))) {
			filteredHolidays = append(filteredHolidays, holiday)
		}
	}

	return filteredHolidays, nil
}

func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
