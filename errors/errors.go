package errors

import (
	"errors"
	"fmt"
)

type KError struct {
	code    int
	message string
	parent  error
}

func New(code int, message string, err ...error) error {
	var parent error

	if len(err) > 0 {
		parent = err[0]
	}

	return KError{
		code:    handleCode(code),
		message: message,
		parent:  parent,
	}
}

func (ke KError) Code() int {
	return ke.code
}

func (ke KError) Message() string {
	return ke.message
}

func (ke KError) Error() string {
	if CurrentSettings.ErrorFormatter == nil {
		return ke.defaultErrorFormatter()
	}
	return CurrentSettings.ErrorFormatter(ke) + ke.parentMessage()
}

func (ke KError) Unwrap() error {
	return ke.parent
}

func (ke KError) parentMessage() string {
	if ke.parent == nil {
		return ""
	}
	return fmt.Sprintf("%s%s", CurrentSettings.Separator, ke.parent.Error())
}

func (ke KError) defaultErrorFormatter() string {
	return fmt.Sprintf("[ERR%d] %s%s", ke.code, ke.message, ke.parentMessage())
}

func Wrap(err error, code int, message string) error {
	return New(code, message, err)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func handleCode(code int) int {
	if code != 0 {
		return code
	}
	return CurrentSettings.DefaultCode
}
