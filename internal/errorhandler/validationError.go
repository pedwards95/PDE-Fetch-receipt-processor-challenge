package errorhandler

import "fmt"

// ValidationError wrapper. If more than one param is input, it treats the first one as the format string, and the rest as inputs
func ValidationError(field string, params ...interface{}) *Error {
	detail := &Detail{
		Field: field,
	}
	errMsg := fmt.Sprintf("This field or request did not pass validation. Field '%s'.", field)
	returnErr := &errorSpecial{
		ErrorMessage: errMsg,
		Detail:       detail,
		HTTPCode:     400,
		MessageType:  "ValidationError",
	}
	if len(params) > 0 {
		infoStr, err := formatParams(params)
		if err != nil {
			return err
		}
		detail.Info = infoStr
		return errorf(returnErr, detail.Info)
	}
	return errorf(returnErr)
}
