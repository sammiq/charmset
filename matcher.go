package matchers

type Matcher interface {
	Description() string
	Match(actual interface{}) error
}

type MatcherType struct {
	expected  string
	matchFunc func(actual interface{}) error
}

func (mt *MatcherType) Description() string {
	return mt.expected
}

func (mt *MatcherType) Match(actual interface{}) error {
	return mt.matchFunc(actual)
}

func NewMatcher(description string, match func(actual interface{}) error) *MatcherType {
	return &MatcherType{description, match}
}
