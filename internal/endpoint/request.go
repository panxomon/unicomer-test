package endpoint

import "time"

type Request struct {
	Date        time.Time `json:"date" xml:"date"`
	Title       string    `json:"title" xml:"title"`
	Type        string    `json:"type" xml:"type"`
	Inalienable bool      `json:"inalienable" xml:"inalienable"`
	Extra       string    `json:"extra" xml:"extra"`
}
