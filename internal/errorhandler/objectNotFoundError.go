package errorhandler

import "fmt"

// ObjectNotFoundError wrapper. If more than one param is input, it treats the first one as the format string, and the rest as inputs
func ObjectNotFoundError(fieldType string, id string, params ...interface{}) *Error {
	detail := &Detail{
		Type: fieldType,
		ID:   id,
	}
	errMsg := fmt.Sprintf("This object doesn't exist. Field '%s'. Id '%s'", fieldType, id)
	returnErr := &errorSpecial{
		ErrorMessage: errMsg,
		Detail:       detail,
		HTTPCode:     404,
		MessageType:  "ObjectNotFoundError",
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
