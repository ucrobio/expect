package expect

func Value[T any](value T) Expector[T] {
	return new(expector[T]).init(value)
}
