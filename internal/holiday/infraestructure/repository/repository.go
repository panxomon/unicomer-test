package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	"net/http"
	"unicomer-test/internal/holiday/domain"
)

type HolidayRepository struct {
	url string
}

func NewHolidayRepository(url string) domain.HolidayRepository {
	return &HolidayRepository{
		url: url,
	}
}

func (h *HolidayRepository) Retrieve(ctx context.Context) ([]domain.Holiday, error) {
	res, err := http.Get(h.url)
	if err != nil {
		log.Ctx(ctx).Err(err).Err(err).Msg("error while get data")
	}

	defer res.Body.Close()

	var result struct {
		Status string           `json:"status"`
		Data   []domain.Holiday `json:"data"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Ctx(ctx).Err(err).Msg("error while decode response")
	}

	// Parse dates
	for i, holiday := range result.Data {
		date, _ := time.Parse("2006-01-02", holiday.Date.Format("2006-01-02"))
		result.Data[i].Date = date
	}

	return result.Data, nil
}
