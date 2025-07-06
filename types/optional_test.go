package types

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type OptionalStruct struct {
	OptionalField Optional[Nullable[int]] `json:"optional_field,omitzero" validate:"omitempty,max=100"`
}

func TestValidateOptional(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterCustomTypeFunc(ValuerCustomTypeFunc, Optional[Nullable[int]]{}, Nullable[int]{})

	var testCases = []struct {
		name        string
		input       OptionalStruct
		shouldError bool
	}{
		{
			name:        "Missing validates",
			input:       OptionalStruct{Empty[Nullable[int]]()},
			shouldError: false,
		},
		{
			name:        "Optional some validates",
			input:       OptionalStruct{WithValue(Some(10))},
			shouldError: false,
		},
		{
			name:        "Optional Nil validates",
			input:       OptionalStruct{WithValue(Nil[int]())},
			shouldError: false,
		},
		{
			name:        "Too large of value fails",
			input:       OptionalStruct{WithValue(Some(1_000))},
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
