package errorhandler

import (
	"fmt"
	"reflect"
)

// Error special error object, implements error
type Error struct {
	Detail        *Detail `json:"detail"`
	ErrorHTTPCode int     `json:"http_code"`
	Text          string  `json:"message"`
	Root          string  `json:"root"`
}

// errorSpecial internal error object, in case you wanted to do special logging, handling, fields, etc
type errorSpecial struct {
	Detail       *Detail
	ErrorMessage string
	HTTPCode     int
	MessageType  string
}

// Error ...
func (e *Error) Error() string {
	return e.Text
}

// AddRootError for when you want to pass up the native error
func (e *Error) AddRootError(err error) *Error {
	if err != nil && len(err.Error()) > 0 {
		e.Root = err.Error()
	} else {
		e.Root = "Nil or empty error passed to root"
	}
	return e
}

// errorf for formatting the individual special errors into the uniform Error
func errorf(err *errorSpecial, text ...string) *Error {
	// user message with optional extra text
	returnMessage := ""
	if len(text) > 0 {
		returnMessage = fmt.Sprintf("%d %s: %s %s", err.HTTPCode, err.MessageType, err.ErrorMessage, text)
	} else {
		returnMessage = fmt.Sprintf("%d %s: %s", err.HTTPCode, err.MessageType, err.ErrorMessage)
	}

	//craft base error
	internalErr := &Error{
		ErrorHTTPCode: err.HTTPCode,
		Text:          returnMessage,
	}

	//for errors that are not base error
	if err.Detail != nil {
		internalErr.Detail = err.Detail
	}

	return internalErr
}

// formatParams helper function for extra string formatting; allows flexibility
func formatParams(params []interface{}) (string, *Error) {
	switch len(params) {
	case 0:
		return "", nil
	case 1:
		if reflect.TypeOf(params[0]).Kind() != reflect.String {
			return "", InternalError("Non-string passed in a format string in errorhandler")
		}
		return params[0].(string), nil
	default:
		if reflect.TypeOf(params[0]).Kind() != reflect.String {
			return "", InternalError("Non-string passed in a format string in errorhandler")
		}
		return fmt.Sprintf(params[0].(string), params[1:]...), nil
	}
}
