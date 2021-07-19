package http

import "time"

type Request struct {
	Code string    `json:"code" validate:"required,len=6"`
	Time time.Time `json:"time" validate:"required"`
}
