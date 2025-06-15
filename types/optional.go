package types

import (
	"fmt"
)

// Optional is a generic type that may or may not contain a non-nil value.
// T is the type of the value.
type Optional[T any] struct {
	value   T
	present bool
}

// Empty returns an empty Optional instance.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// WithValue returns an Optional describing the given value.
// If the provided value is considered "zero" for its type (e.g. nil for pointers/interfaces,
// 0 for numbers, "" for strings), it's still considered present.
// Use Filter or other methods if you need to treat specific zero values as absent.
func WithValue[T any](value T) Optional[T] {
	return Optional[T]{
		value:   value,
		present: true,
	}
}

// HasValue is an alias for IsPresent.
// Returns true if there is a value present, otherwise false.
func (o Optional[T]) HasValue() bool {
	return o.present
}

// Get returns the value if present and a boolean indicating whether the value was present.
// This is the idiomatic Go way to access a potentially missing value.
func (o Optional[T]) Get() (T, bool) {
	if !o.present {
		var zero T
		return zero, false
	}
	return o.value, true
}

// MustGet returns the value if present. It panics if the value is not present.
// Use with caution, typically when you are certain the value is present.
func (o Optional[T]) MustGet() T {
	if !o.present {
		var zero T // To get type name for panic message
		panic(fmt.Sprintf("optional: MustGet() called on an empty Optional[%T]", zero))
	}
	return o.value
}

// OrElse returns the value if present, otherwise returns other.
func (o Optional[T]) OrElse(other T) T {
	if o.present {
		return o.value
	}
	return other
}

// String returns a string representation of the Optional.
// For example, "Optional[int](10)" or "Optional[string](empty)".
func (o Optional[T]) String() string {
	if o.present {
		return fmt.Sprintf("Optional[%T](%v)", o.value, o.value)
	}
	// To get the type name T even when empty
	var zero T
	return fmt.Sprintf("Optional[%T](empty)", zero)
}
