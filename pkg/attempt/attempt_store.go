package attempt

import (
	"fmt"
	"sync"
	"time"
)

type AttemptStore interface {
	Get(string) (*Attempt, error)
	Put(*Attempt) error
	Del(string) error
}

type InMemoryAttemptStore struct {
	locker sync.RWMutex
	db     map[string]*Attempt
}

func NewInMemoryAttemptStore() *InMemoryAttemptStore {
	return &InMemoryAttemptStore{
		locker: sync.RWMutex{},
		db:     make(map[string]*Attempt),
	}
}

func (as *InMemoryAttemptStore) Get(ip string) (*Attempt, error) {
	attempt, ok := as.db[ip]

	if !ok {
		return nil, fmt.Errorf("Attempt for ip %s does not exist", ip)
	}

	return attempt, nil
}

func (as *InMemoryAttemptStore) Put(attempt *Attempt) error {
	as.locker.Lock()

	defer as.locker.Unlock()

	curr, err := as.Get(attempt.Ip)

	if err != nil {
		attempt.Quantity = 1
		attempt.CreatedAt = uint32(time.Now().UnixNano())
		as.db[attempt.Ip] = attempt

		return nil
	}
	curr.Quantity = curr.Quantity + 1
	as.db[curr.Ip] = curr

	return nil

}

func (as *InMemoryAttemptStore) Del(ip string) error {
	as.locker.Lock()
	defer as.locker.Unlock()

	delete(as.db, ip)

	return nil
}
