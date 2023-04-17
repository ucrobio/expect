package expect

import "fmt"

type matcherBeError struct {
}

func (it matcherBeError) init() matcherBeError {
	return it
}

func (it matcherBeError) PositiveMatch(value error) bool {
	return value != nil
}

func (it matcherBeError) NegativeMatch(value error) bool {
	return value == nil
}

func (it matcherBeError) PositiveError(value error) error {
	return fmt.Errorf("should be error")
}

func (it matcherBeError) NegativeError(value error) error {
	return fmt.Errorf("shouldn't be error: %w", value)
}
