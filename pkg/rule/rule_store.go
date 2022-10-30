package rule

import "fmt"

type RuleStore interface {
	Put(*Rule) error
	Get(string) (*Rule, error)
}

type InMemoryRuleStore struct {
	db map[string]*Rule
}

func NewInMemoryRuleStore() *InMemoryRuleStore {
	return &InMemoryRuleStore{
		db: make(map[string]*Rule),
	}
}

func (r *InMemoryRuleStore) Put(rule *Rule) error {
	r.db[rule.Region] = rule
	return nil
}

func (r *InMemoryRuleStore) Get(region string) (*Rule, error) {
	rule, ok := r.db[region]
	if !ok {
		return nil, fmt.Errorf("named rule:%s does not exist", region)
	}

	return rule, nil
}
