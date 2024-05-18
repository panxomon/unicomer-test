package bootstrap

import (
	"fmt"
	holiday "unicomer-test/internal/holiday/application"
	holidayFind "unicomer-test/internal/holiday/application/find"
	holidayrepo "unicomer-test/internal/holiday/infraestructure/repository"
)

type Bootstrap struct {
	Holidays *holiday.App
}

func LoadComponents(urlHolidaysServer string) (*Bootstrap, error) {
	holidayRepository, err := holidayrepo.MakeHolidayRepository(urlHolidaysServer)
	if err != nil {
		return nil, fmt.Errorf("error while getting data: %v", err)
	}

	holidayService := holidayFind.NewHolidayFinder(holidayRepository)
	holidayQuery := holidayFind.NewHolidayFindQueryHandler(holidayService)

	holidaysApp := holiday.NewApp(holiday.Queries{FindCodebase: holidayQuery})

	return &Bootstrap{
		Holidays: holidaysApp,
	}, nil
}
