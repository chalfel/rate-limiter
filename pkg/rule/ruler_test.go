package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRulerPutGet(t *testing.T) {
	ruler := NewRuler(createInMemoryRuleStore(t))
	region := "test"
	rule := &Rule{
		Region: region,
	}

	err := ruler.PutRule(rule)

	assert.Nil(t, err)

	sut, err := ruler.GetRule(region)

	assert.Nil(t, err)
	assert.Equal(t, sut, rule)

}
