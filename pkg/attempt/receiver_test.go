package attempt

import (
	"testing"
	"time"

	"github.com/chalfel/rate-limiter/pkg/rule"
	"github.com/stretchr/testify/assert"
)

func TestReceiverTryShouldReturnNoError(t *testing.T) {
	ruleStore := rule.NewInMemoryRuleStore()
	attemptStore := createInMemoryAttemptStore(t)
	ruleStore.Put(&rule.Rule{Region: "test", Limit: 10, Duration: 1000})
	rc := NewReceiver(attemptStore, *rule.NewRuler(ruleStore))

	assert.NotNil(t, rc)
	attempt := &Attempt{
		Ip:     "test",
		Region: "test",
	}
	err := rc.Try(attempt.Ip, attempt.Region)

	assert.Nil(t, err)
	createdAttempt, err := attemptStore.Get(attempt.Ip)

	assert.Nil(t, err)
	assert.Equal(t, createdAttempt.Ip, attempt.Ip)
}

func TestReceiverTryShouldReturnErrorIfAttempSurpassLimit(t *testing.T) {
	ruleStore := rule.NewInMemoryRuleStore()
	attemptStore := createInMemoryAttemptStore(t)
	ruleStore.Put(&rule.Rule{Region: "test", Limit: 1, Duration: uint32(time.Second * 1)})
	rc := NewReceiver(attemptStore, *rule.NewRuler(ruleStore))

	assert.NotNil(t, rc)
	attempt := &Attempt{
		Ip:     "test",
		Region: "test",
	}
	err := rc.Try(attempt.Ip, attempt.Region)

	assert.Nil(t, err)

	err = rc.Try(attempt.Ip, attempt.Region)

	assert.NotNil(t, err)
}

func TestReceiverTryShouldReturnNoErrorAfterRuleDurationIsFinished(t *testing.T) {
	ruleStore := rule.NewInMemoryRuleStore()
	attemptStore := createInMemoryAttemptStore(t)
	ruleStore.Put(&rule.Rule{Region: "test", Limit: uint32(1), Duration: uint32(time.Nanosecond * 1)})
	rc := NewReceiver(attemptStore, *rule.NewRuler(ruleStore))

	assert.NotNil(t, rc)
	attempt := &Attempt{
		Ip:     "test",
		Region: "test",
	}
	err := rc.Try(attempt.Ip, attempt.Region)

	assert.Nil(t, err)

	time.Sleep(100 * time.Millisecond)
	err = rc.Try(attempt.Ip, attempt.Region)

	assert.Nil(t, err)
}
