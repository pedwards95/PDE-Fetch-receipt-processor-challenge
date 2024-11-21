package errorhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	detail := &Detail{
		Field: "banana",
	}
	err := &errorSpecial{
		ErrorMessage: "I can't make coffee!",
		HTTPCode:     419,
		MessageType:  "TeapotError",
		Detail:       detail,
	}
	expectedError := &Error{
		Text:          "419 TeapotError: I can't make coffee!",
		ErrorHTTPCode: 419,
		Detail:        detail,
	}
	errRes := errorf(err)
	assert.IsType(t, expectedError, errRes)
	assert.Equal(t, expectedError.Text, errRes.Text)

	errRes = errorf(err, "Try your kitchen!")
	expectedError.Text = expectedError.Text + " [Try your kitchen!]"
	assert.Equal(t, expectedError.Text, errRes.Text)
}

func TestError(t *testing.T) {
	err := &errorSpecial{
		ErrorMessage: "I can't make coffee!",
		HTTPCode:     419,
		MessageType:  "TeapotError",
	}
	errRes := errorf(err)
	assert.Equal(t, "419 TeapotError: I can't make coffee!", errRes.Error())
}
