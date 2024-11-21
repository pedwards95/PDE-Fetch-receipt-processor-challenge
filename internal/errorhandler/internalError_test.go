package errorhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {
	err := InternalError()
	assert.Equal(t, "500 InternalServerError: Internal server error.", err.Text)
	err = InternalError("Yep. Everything is broken.")
	assert.Equal(t, "500 InternalServerError: Internal server error. [Yep. Everything is broken.]", err.Text)
}
