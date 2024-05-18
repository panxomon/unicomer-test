package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
	"time"
	holiday "unicomer-test/internal/holiday/application"
	"unicomer-test/internal/holiday/application/find"
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

func (h *Holidays) GenerateProject(c *gin.Context) error {
	ctx := c.Request.Context()
	
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	holidays, err := h.holidayApp.Queries.FindCodebase.Handle(ctx, find.HolidayFindQuery{HolidayType: "", Date: time.Time{}})
	if err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Invalid request payload"})
	}

	c.JSON(http.StatusOK, gin.H{"holidays": holidays})
	return nil
}
