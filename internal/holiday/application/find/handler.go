package find

import (
	"context"
	"unicomer-test/internal/cqrs"
	"unicomer-test/internal/holiday/domain"
)

type HolidayFindQueryHandler cqrs.QueryHandler[HolidayFindQuery, *domain.Holiday]
type holidayFindQueryHandler struct {
	finder domain.HolidayFindQuery
}

func NewHolidayFindQueryHandler(query domain.HolidayFindQuery) HolidayFindQueryHandler {
	return &holidayFindQueryHandler{finder: query}
}

func (h *holidayFindQueryHandler) Handle(ctx context.Context, query HolidayFindQuery) (*domain.Holiday, error) {
	return h.finder.Execute(ctx, query.HolidayType, query.Date)
}
