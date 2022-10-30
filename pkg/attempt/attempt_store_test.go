package attempt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createInMemoryAttemptStore(t *testing.T) *InMemoryAttemptStore {

	as := NewInMemoryAttemptStore()

	assert.NotNil(t, as)

	return as

}
func TestAttemptStoreGetPut(t *testing.T) {
	as := createInMemoryAttemptStore(t)

	attempt := &Attempt{Ip: "test"}
	err := as.Put(attempt)

	assert.Nil(t, err)

	currentAttemp, err := as.Get(attempt.Ip)

	assert.Nil(t, err)

	assert.Equal(t, currentAttemp, attempt)
}

func TestAttemptStoreGetShouldReturnErrorIfDoesNotExist(t *testing.T) {
	as := createInMemoryAttemptStore(t)

	attempt := Attempt{Ip: "test"}

	currentAttemp, err := as.Get(attempt.Ip)

	assert.NotNil(t, err)

	assert.Nil(t, currentAttemp)
}

func TestAttemptStoreShouldUpdateQuantityWhenAlreadyExist(t *testing.T) {
	as := createInMemoryAttemptStore(t)

	attempt := &Attempt{Ip: "test"}
	err := as.Put(attempt)

	assert.Nil(t, err)

	err = as.Put(attempt)

	assert.Nil(t, err)

	currentAttemp, err := as.Get(attempt.Ip)

	assert.Nil(t, err)

	assert.Equal(t, currentAttemp.Quantity, uint32(2))
}

func TestAttemptStoreDel(t *testing.T) {
	as := createInMemoryAttemptStore(t)

	attempt := &Attempt{Ip: "test"}
	err := as.Put(attempt)

	assert.Nil(t, err)

	err = as.Del(attempt.Ip)

	assert.Nil(t, err)

	currentAttemp, err := as.Get(attempt.Ip)

	assert.Nil(t, currentAttemp)
	assert.NotNil(t, err)
}
