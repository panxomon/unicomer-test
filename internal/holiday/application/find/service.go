package find

import (
	"context"
	"unicomer-test/internal/holiday/domain"

	"github.com/rs/zerolog/log"
)

type HolidayFinderService struct {
	repository domain.HolidayRepository
}

func NewHolidayFinder(repository domain.HolidayRepository) domain.HolidayFindQuery {
	return &HolidayFinderService{
		repository: repository,
	}
}

func (c *HolidayFinderService) Execute(ctx context.Context) ([]domain.Holiday, error) {
	holidaysResp, err := c.repository.Retrieve(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Str("Holiday", "find").Msg("failed executing")
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
