package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type NullableStruct struct {
	NullableField Nullable[int] `json:"nullable_int,omitempty" validate:"required,max=100"`
}

func TestJSONEncoding(t *testing.T) {
	testCases := []struct {
		name  string
		input NullableStruct
		want  string
	}{
		{
			name:  "marshal not null",
			input: NullableStruct{NullableField: Some(10)},
			want:  "{\"nullable_int\":10}",
		},
		{
			name:  "marshal null",
			input: NullableStruct{NullableField: Nil[int]()},
			want:  "{\"nullable_int\":null}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got, err := json.Marshal(tc.input)
			assert.NoError(t, err)

			assert.Equal(t, tc.want, string(got))
		})
	}
}

func TestJSONDecoding(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  NullableStruct
	}{
		{
			name:  "unmarshal not null",
			input: "{\"nullable_int\":10}",
			want:  NullableStruct{NullableField: Some(10)},
		},
		{
			name:  "marshal null",
			input: "{\"nullable_int\":null}",
			want:  NullableStruct{NullableField: Nil[int]()},
		},
		{
			name:  "unmarshal missing",
			input: "{}",
			want:  NullableStruct{NullableField: Nullable[int]{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got NullableStruct
			err := json.Unmarshal([]byte(tc.input), &got)
			assert.NoError(t, err)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateWorks(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterCustomTypeFunc(ValuerCustomTypeFunc, Nullable[any]{})

	v := NullableStruct{Some(10)}
	err := validate.Struct(v)
	assert.NoError(t, err)
}

func TestValidateRequiredWorks(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterCustomTypeFunc(ValuerCustomTypeFunc, Nullable[int]{})

	v := NullableStruct{}
	err := validate.Struct(v)
	assert.Error(t, err)
}

func TestValidateTypeConstraintsWorks(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterCustomTypeFunc(ValuerCustomTypeFunc, Nullable[int]{})

	v := NullableStruct{Some(1000)}
	err := validate.Struct(v)
	a := fmt.Sprint(err.Error())
	_ = a
	assert.Error(t, err)
}
