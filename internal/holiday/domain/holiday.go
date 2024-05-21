package domain

import (
	"context"
	"time"
)

type Holiday struct {
	Date        time.Time `json:"date" xml:"date"`
	Title       string    `json:"title" xml:"title"`
	Type        string    `json:"type" xml:"type"`
	Inalienable bool      `json:"inalienable" xml:"inalienable"`
	Extra       string    `json:"extra" xml:"extra"`
}

type HolidayResponse struct {
	Status string    `json:"status" xml:"status"`
	Data   []Holiday `json:"data" xml:"data"`
}

type HolidayFindQuery interface {
	Execute(ctx context.Context) ([]Holiday, error)
}

type HolidayRepository interface {
	Retrieve(ctx context.Context) (*HolidayResponse, error)
}
