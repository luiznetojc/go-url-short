package model

import "time"

type Url struct {
	Id string `json:"id"`
	UserId string
	OriginalUrl string
	ShortUrl string
	CreatedAt time.Time
	ExpiredAt time.Time
}