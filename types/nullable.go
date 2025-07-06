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
	value  *T
	isNull bool
}

func Some[T any](value T) Nullable[T] {
	return Nullable[T]{
		value:  &value,
		isNull: false,
	}
}

func Nil[T any]() Nullable[T] {
	var zero T
	return Nullable[T]{value: &zero, isNull: true}
}

func (n Nullable[T]) HasValue() bool {
	return !n.isNull
}

func (n Nullable[T]) Unwrap() (T, error) {
	if n.HasValue() {
		return *n.value, nil
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
	if !n.isNull {
		return fmt.Sprintf("Nullable[%T](%v)", n.value, n.value)
	}
	// To get the type name T even when empty
	var zero T
	return fmt.Sprintf("Nullable[%T](null)", zero)
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.isNull {
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
		n.value = value
		n.isNull = false
	} else {
		var zero T
		n.value = &zero
		n.isNull = true
	}
	return nil
}

func (n Nullable[T]) Value() (Value, error) {
	if !n.isNull {
		return n.value, nil
	}

	var zero T
	return &zero, nil
}
