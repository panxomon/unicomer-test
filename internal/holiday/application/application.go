package application

import "unicomer-test/internal/holiday/application/find"

type App struct {
	Queries Queries
}

func NewApp(queries Queries) *App {
	return &App{Queries: queries}
}

type Queries struct {
	FindCodebase find.HolidayFindQueryHandler
}
