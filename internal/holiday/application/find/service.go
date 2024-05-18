package find

import (
	"context"
	"github.com/rs/zerolog/log"
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

func (c *HolidayFinderService) Execute(ctx context.Context, programmingLanguage, frameworkName, frameworkVersion string) (*domain.Holiday, error) {
	holidays, err := c.repository.Retrieve(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Str("codebase", "find").Msg("failed executing")
		return nil, err
	}

	return holidays, nil
}
