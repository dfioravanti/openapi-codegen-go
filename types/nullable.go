package types

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	CannotUnwrapNilErr = errors.New("cannot unwrap nil")
)

type Nullable[T any] struct {
	value T
	null  bool
}

func Some[T any](value T) Nullable[T] {
	return Nullable[T]{
		value: value,
		null:  false,
	}
}

func Nil[T any]() Nullable[T] {
	return Nullable[T]{null: true}
}

func (n Nullable[T]) HasValue() bool {
	return !n.null
}

func (n Nullable[T]) Unwrap() (T, error) {
	if n.HasValue() {
		return n.value, nil
	}
	var zero T
	return zero, CannotUnwrapNilErr
}

func (n Nullable[T]) MustUnwrap() T {
	v, err := n.Unwrap()
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a string representation of the Nullable.
// For example, "Nullable[int](10)" or "Nullable[string](null)".
func (n Nullable[T]) String() string {
	if !n.null {
		return fmt.Sprintf("Nullable[%T](%v)", n.value, n.value)
	}
	// To get the type name T even when empty
	var zero T
	return fmt.Sprintf("Nullable[%T](null)", zero)
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.null {
		return []byte("null"), nil
	}

	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	var value *T
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value != nil {
		n.value = *value
		n.null = false
	} else {
		n.null = true
	}
	return nil
}

func (n Nullable[T]) Value() (Value, error) {
	return n.value, nil
}
