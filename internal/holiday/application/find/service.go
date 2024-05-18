package find

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
	"unicomer-test/internal/holiday/domain"
)

type HolidayFinderService struct {
	repository domain.HolidayRepository
}

func NewHolidayFinder(repository domain.HolidayRepository) domain.HolidayFindQuery {
	return &HolidayFinderService{
		repository: repository,
	}
}

func (c *HolidayFinderService) Execute(ctx context.Context, holidayType string, date time.Time) ([]domain.Holiday, error) {
	holidaysResp, err := c.repository.Retrieve(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Str("codebase", "find").Msg("failed executing")
		return nil, err
	}

	var holidays []domain.Holiday

	for _, holidayData := range holidaysResp.Data {
		holiday := domain.Holiday{
			Date:        holidayData.Date,
			Title:       holidayData.Title,
			Type:        holidayData.Type,
			Inalienable: holidayData.Inalienable,
			Extra:       holidayData.Extra,
		}
		holidays = append(holidays, holiday)
	}

	return holidays, nil
}
