package expect

type (
	Matcher[T any] interface {
		PositiveMatch(value T) bool
		PositiveError(value T) error
		NegativeMatch(value T) bool
		NegativeError(value T) error
	}
)
