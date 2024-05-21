package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
	"unicomer-test/internal/holiday/domain"
)

type holidayRepository struct {
	data *domain.HolidayResponse
}

func MakeHolidayRepository(ctx context.Context, url string) (domain.HolidayRepository, error) {

	res, err := http.Get(url)
	if err != nil {
		log.Ctx(ctx).Err(err).Str("project", "holiday").Msg("error while getting data")
		return nil, fmt.Errorf("error while getting data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("non-OK status code received: %d", res.StatusCode)
		log.Ctx(ctx).Err(err).Str("project", "holiday").Int("status_code", res.StatusCode).Msg("received non-OK status code")
		return nil, fmt.Errorf("non-OK status code received: %d", res.StatusCode)
	}

	var tempData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&tempData); err != nil {
		log.Ctx(ctx).Err(err).Str("project", "holiday").Msg("error while decoding response")
		return nil, fmt.Errorf("error while decoding response: %v", err)
	}

	data := domain.HolidayResponse{
		Status: tempData["status"].(string),
	}
	holidays := make([]domain.Holiday, 0)
	for _, holidayData := range tempData["data"].([]interface{}) {
		holidayMap := holidayData.(map[string]interface{})
		DateStr := holidayMap["date"].(string)
		Date, err := time.Parse("2006-01-02", DateStr)
		if err != nil {
			log.Ctx(ctx).Err(err).Str("project", "holiday").Str("date_string", DateStr).Msg("error while parsing date")
			break
		}
		holiday := domain.Holiday{
			Date:        Date,
			Title:       holidayMap["title"].(string),
			Type:        holidayMap["type"].(string),
			Inalienable: holidayMap["inalienable"].(bool),
			Extra:       holidayMap["extra"].(string),
		}
		holidays = append(holidays, holiday)
	}
	data.Data = holidays

	repo := &holidayRepository{
		data: &data,
	}

	return repo, nil
}

func (r *holidayRepository) Retrieve(ctx context.Context) (*domain.HolidayResponse, error) {
	return r.data, nil
}
