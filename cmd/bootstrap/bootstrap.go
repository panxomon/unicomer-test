package bootstrap

import (
	holiday "unicomer-test/internal/holiday/application"
	holidayFind "unicomer-test/internal/holiday/application/find"
	holidayrepo "unicomer-test/internal/holiday/infraestructure/repository"
)

type Bootstrap struct {
	Holidays *holiday.App
}

func LoadComponents(urlHolidaysServer string) *Bootstrap {
	holidayRepository := holidayrepo.NewHolidayRepository(urlHolidaysServer)
	holidayService := holidayFind.NewHolidayFinder(holidayRepository)
	holidayQuery := holidayFind.NewHolidayFindQueryHandler(holidayService)

	holidaysApp := holiday.NewApp(holidayQuery)

	return &Bootstrap{
		Holidays: holidaysApp,
	}
}
