package attempt

import (
	"fmt"
	"time"

	"github.com/chalfel/rate-limiter/pkg/rule"
	"github.com/sirupsen/logrus"
)

type Receiver struct {
	store AttemptStore
	ruler rule.Ruler
}

func NewReceiver(store AttemptStore, ruler rule.Ruler) *Receiver {
	return &Receiver{
		store: store,
		ruler: ruler,
	}
}

func (r *Receiver) Try(ip string, region string) error {
	rule, err := r.ruler.GetRule(region)

	if err != nil {
		logrus.WithFields(logrus.Fields{"Ip": ip, "Region": region}).Infoln("rule does not exist")
		panic(err)
	}

	currentAttempt, err := r.store.Get(ip)

	if err != nil {
		attempt := &Attempt{Ip: ip, Region: region, Quantity: 1, CreatedAt: uint32(time.Now().UnixNano())}
		r.store.Put(attempt)

		return nil
	}

	if currentAttempt.Quantity < rule.Limit {
		r.store.Put(currentAttempt)

		return nil
	}

	if currentAttempt.CreatedAt+rule.Duration <= uint32(time.Now().UnixNano()) {
		r.store.Del(ip)

		attempt := &Attempt{Ip: ip, Region: region, Quantity: 1, CreatedAt: uint32(time.Now().UnixNano())}

		r.store.Put(attempt)
		return nil
	}

	logrus.WithFields(logrus.Fields{"Ip": ip, "Region": region}).Infoln("limit reached")

	return fmt.Errorf("limit reached for ip %s", ip)
}
