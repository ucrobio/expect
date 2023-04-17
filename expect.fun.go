package expect

func Value[T any](value T) Expector[T] { return new(expector[T]).init(value) }
func Error(err error) Expector[error]  { return Value(err) }

func BeEq[T comparable](value T) Matcher[T] { return new(matcherBeEq[T]).init(value) }
func BeError() Matcher[error]               { return new(matcherBeError).init() }
