package rule

type Rule struct {
	Region   string
	Limit    uint32
	Duration uint32
}

type Ruler struct {
	store RuleStore
}

func NewRuler(
	store RuleStore,
) *Ruler {
	store.Put(&Rule{Region: "123", Limit: 2, Duration: uint32(1000)})
	return &Ruler{
		store: store,
	}
}

func (r *Ruler) GetRule(name string) (*Rule, error) {
	return r.store.Get(name)
}

func (r *Ruler) PutRule(rule *Rule) error {
	return r.store.Put(rule)
}
