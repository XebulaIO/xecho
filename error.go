package xecho

import (
	"fmt"
	"net/http"
	"time"
)

type Error struct {
	Code         string `json:"code"`
	Detail       string `json:"detail"`
	TrackID      string `json:"track_id"`
	Timestamp    int64  `json:"timestamp"`
	responseCode int
}

func NewError(c XContext) *Error {
	return &Error{
		TrackID:      c.ID(),
		Timestamp:    time.Now().Unix(),
		responseCode: http.StatusInternalServerError,
	}
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] -> %s: %s", e.TrackID, e.Code, e.Detail)
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) WithCode(code string) *Error {
	e.Code = code
	return e
}

func (e *Error) WithDetail(detail error) *Error {
	e.Detail = detail.Error()
	return e
}

func (e *Error) WithResponseCode(code int) *Error {
	e.responseCode = code
	return e
}
