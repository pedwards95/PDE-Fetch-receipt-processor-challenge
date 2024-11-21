package errorhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	err := ValidationError("MyField")
	assert.Equal(t, "400 ValidationError: This field or request did not pass validation. Field 'MyField'.", err.Text)
	assert.Equal(t, "MyField", err.Detail.Field)
	err = ValidationError("MyField", "Just no. %s", "Bad data.")
	assert.Equal(t, "400 ValidationError: This field or request did not pass validation. Field 'MyField'. [Just no. Bad data.]", err.Text)
}
