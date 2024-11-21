package errorhandler

// InternalError wrapper. If more than one param is input, it treats the first one as the format string, and the rest as inputs
func InternalError(params ...interface{}) *Error {
	detail := &Detail{}
	errMsg := "Internal server error."
	returnErr := &errorSpecial{
		ErrorMessage: errMsg,
		Detail:       detail,
		HTTPCode:     500,
		MessageType:  "InternalServerError",
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
