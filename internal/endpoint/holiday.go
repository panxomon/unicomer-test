package endpoint

import (
	"context"
	"net/http"
	"strings"
	"time"
	holiday "unicomer-test/internal/holiday/application"
	"unicomer-test/internal/holiday/application/find"
	"unicomer-test/internal/holiday/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Holidays struct {
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
	acceptHeader := c.GetHeader("Content-Type")

	holidays, err := h.FilterHolidays(ctx, typeFilter, startDate, endDate)
	if err != nil {
		log.Ctx(ctx).Err(err).Str("project", "holiday").Msg("Failed to filter holidays")
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

	allHolidays, err := h.holidayApp.Queries.FindHoliday.Handle(ctx, find.HolidayFindQuery{HolidayType: "", Date: time.Time{}})
	if err != nil {
		log.Ctx(ctx).Err(err).Str("project", "holiday").Msg("Error retrieving holidays from repository")
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
	log.Ctx(ctx).Info().Str("project", "holiday").Int("filtered_holidays_count", len(filteredHolidays)).Msg("Filtered holidays successfully")
	return filteredHolidays, nil
}

func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Error().Err(err).Str("project", "holiday").Str("date_string", dateStr).Msg("Error parsing date")
		return time.Time{}
	}
	return date
}
