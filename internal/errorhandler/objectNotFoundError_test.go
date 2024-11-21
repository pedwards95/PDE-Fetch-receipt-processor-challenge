package errorhandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectNotFoundError(t *testing.T) {
	err := ObjectNotFoundError("MyField", "1")
	assert.Equal(t, "404 ObjectNotFoundError: This object doesn't exist. Field 'MyField'. Id '1'", err.Text)
	assert.Equal(t, "MyField", err.Detail.Type)
	err = ObjectNotFoundError("MyField", "3PO", "Where could he be?")
	assert.Equal(t, "404 ObjectNotFoundError: This object doesn't exist. Field 'MyField'. Id '3PO' [Where could he be?]", err.Text)
}
