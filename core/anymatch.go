package core

import "github.com/sammiq/matchers"

func anyMatch(matchers []matchers.Matcher, actual interface{}) (err error) {
	for _, matcher := range matchers {
		err = matcher.Match(actual)
		if err == nil {
			break
		}
	}
	return err
}
