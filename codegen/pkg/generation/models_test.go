package generation

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/dfioravanti/openapi-codegen-go/pkg/models"
	"github.com/dfioravanti/openapi-codegen-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateModel(t *testing.T) {
	testCases := []struct {
		name  string
		input models.Model
		want  string
		stop  bool
	}{
		{
			name: "model with only builtin types no comments",
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name:          "IntField",
						GoTypePackage: "models",
						GoType:        "int",
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"type BuiltInModel struct {",
				"	IntField int `json:\"int_field\"`",
				"}",
				"",
			),
		},
		{
			name: "model with only builtin types with multiline comments",
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name: "IntField",
						Comment: utils.Multiline(
							"IntField represents an attribute of type int",
							"which can be used for many things",
						),
						GoTypePackage: "models",
						GoType:        "int",
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"type BuiltInModel struct {",
				"	/*",
				"	   IntField represents an attribute of type int",
				"	   which can be used for many things",
				"	*/",
				"	IntField int `json:\"int_field\"`",
				"}",
				"",
			),
		},
		{
			name: "model with only builtin types with single line comments",
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name:          "IntField",
						Comment:       "IntField represents an attribute of type int",
						GoTypePackage: "models",
						GoType:        "int",
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"type BuiltInModel struct {",
				"	// IntField represents an attribute of type int",
				"	IntField int `json:\"int_field\"`",
				"}",
				"",
			),
		},
		{
			name: "model optional only builtin",
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name:          "IntField",
						GoTypePackage: "models",
						GoType:        "int",
						Optional:      true,
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"import types \"github.com/dfioravanti/openapi-codegen-go/types\"",
				"",
				"type BuiltInModel struct {",
				"	IntField types.Optional[int] `json:\"int_field\"`",
				"}",
				"",
			),
		},
		{
			name: "model nullable only builtin",
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name:          "IntField",
						GoTypePackage: "models",
						GoType:        "int",
						Nullable:      true,
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"import types \"github.com/dfioravanti/openapi-codegen-go/types\"",
				"",
				"type BuiltInModel struct {",
				"	IntField types.Nullable[int] `json:\"int_field\"`",
				"}",
				"",
			),
		},
		{
			name: "model optional and nullable only builtin",
			stop: true,
			input: models.Model{
				Name: "BuiltInModel",
				Fields: []models.Field{
					{
						Name:          "IntField",
						GoTypePackage: "models",
						GoType:        "int",
						Optional:      true,
						Nullable:      true,
						Annotations: map[string]string{
							"json": "int_field",
						},
					},
				},
			},
			want: utils.Multiline(
				"package models",
				"",
				"import types \"github.com/dfioravanti/openapi-codegen-go/types\"",
				"",
				"type BuiltInModel struct {",
				"	IntField types.Optional[types.Nullable[int]] `json:\"int_field\"`",
				"}",
				"",
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := jen.NewFilePath("models")
			f = GenerateModel(f, tc.input)
			got := fmt.Sprintf("%#v", f)

			assert.Equal(t, tc.want, got)
		})
	}
}
