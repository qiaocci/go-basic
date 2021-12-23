package myerror

import "fmt"

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	text string
}

func (e *errorString) Error() string {
	return e.text
}

func Errorf(format string, args ...interface{}) error {
	return New(fmt.Sprintf(format, args...))
}
