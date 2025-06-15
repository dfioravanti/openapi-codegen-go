package types

import "reflect"

type Value any

type Valuer interface {
	Value() (Value, error)
}

func ValuerCustomTypeFunc(field reflect.Value) interface{} {
	if value, ok := field.Interface().(Valuer); ok {
		value, _ := value.Value()
		return value
	}
	return nil
}
