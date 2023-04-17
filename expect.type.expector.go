package expect

type (
	Expector[T any] interface {
		To(matcher Matcher[T]) error
		NotTo(matcher Matcher[T]) error
	}

	expector[T any] struct {
		value T
	}
)

func (it expector[T]) init(value T) expector[T] {
	it.value = value

	return it
}

func (it expector[T]) To(matcher Matcher[T]) error {
	if !matcher.PositiveMatch(it.value) {
		return matcher.PositiveError(it.value)
	}

	return nil
}

func (it expector[T]) NotTo(matcher Matcher[T]) error {
	if !matcher.NegativeMatch(it.value) {
		return matcher.NegativeError(it.value)
	}

	return nil
}
