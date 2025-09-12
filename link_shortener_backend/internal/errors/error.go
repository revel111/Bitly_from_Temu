package errors

import "time"

type HttpError struct {
	Code      int
	Msg       string
	Timestamp time.Time
}

func (e *HttpError) Error() string {
	return e.Msg
}

func NewHttpError(code int, msg string) *HttpError {
	httpError := new(HttpError)
	httpError.Code = code
	httpError.Msg = msg
	httpError.Timestamp = time.Now()
	return httpError
}
