package expect

import "fmt"

type matcherBeEq[T comparable] struct {
	value T
}

func (it matcherBeEq[T]) init(value T) matcherBeEq[T] {
	it.value = value
	return it
}

func (it matcherBeEq[T]) PositiveMatch(expected T) bool {
	return expected == it.value
}

func (it matcherBeEq[T]) NegativeMatch(expected T) bool {
	return expected != it.value
}

func (it matcherBeEq[T]) PositiveError(expected T) error {
	return fmt.Errorf("should be equal: %v != %v", expected, it.value)
}

func (it matcherBeEq[T]) NegativeError(expected T) error {
	return fmt.Errorf("shouldn't be equal: %v == %v", expected, it.value)
}
