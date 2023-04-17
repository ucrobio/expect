package expect

func Value[T any](value T) Expector[T] { return new(expector[T]).init(value) }
func Error(err error) Expector[error]  { return Value(err) }

func BeError() Matcher[error] { return new(matcherBeError).Init() }
