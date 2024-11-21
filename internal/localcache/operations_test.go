package localcache

import (
	"testing"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestObject struct {
	myNum           int64
	myString        string
	myComplexObject *TestObjecter
}

type TestObjecter struct {
	myBool  bool
	myArray []int
}

func TestLocalCache(t *testing.T) {
	insideObj := &TestObjecter{
		myBool:  true,
		myArray: []int{1, 2, 3},
	}
	outsideObj := &TestObject{
		myNum:           11,
		myString:        "no",
		myComplexObject: insideObj,
	}

	lg := logger.New()

	cache, stop, err := New(lg)
	defer stop()
	assert.NoError(t, err)

	id := uuid.New()
	cache.Add(id, outsideObj)

	returnObj := cache.Get(id).(*TestObject)
	assert.Equal(t, outsideObj, returnObj)
	assert.Equal(t, outsideObj.myComplexObject, returnObj.myComplexObject)
}
