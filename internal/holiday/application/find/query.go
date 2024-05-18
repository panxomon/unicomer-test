package find

import (
	"time"
)

type HolidayFindQuery struct {
	HolidayType string    `json:"type"`
	Date        time.Time `json:"date"`
}
