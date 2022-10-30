package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createInMemoryRuleStore(t *testing.T) *InMemoryRuleStore {
	rs := NewInMemoryRuleStore()

	assert.NotNil(t, rs)

	return rs
}
func TestInMemoryRuleStorePutGet(t *testing.T) {
	rs := createInMemoryRuleStore(t)

	region := "test"
	err := rs.Put(&Rule{Region: region})

	assert.Nil(t, err)

	rule, err := rs.Get(region)
	assert.Nil(t, err)

	assert.Equal(t, rule.Region, region)
}

func TestInMemoryRuleStoreShouldReturnErrorIfNoRuleFound(t *testing.T) {
	rs := createInMemoryRuleStore(t)

	ruleName := "test"

	rule, err := rs.Get(ruleName)

	assert.NotNil(t, err)
	assert.Nil(t, rule)
}
