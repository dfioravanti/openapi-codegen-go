package generation

import (
	"github.com/dave/jennifer/jen"

	"github.com/dfioravanti/openapi-codegen-go/pkg/models"
)

const (
	// The name used for the package that contains all the models.
	//
	// Important: we also use this name to generate built in names
	// since jennifer does not render this name as part of a type,
	// so Qual(ModelsPackageName, "int") becomes int
	ModelsPackageName = "models"
)

func GenerateModel(modelFile *jen.File, model models.Model) *jen.File {
	t := jen.Type().Id(model.Name)

	var items []jen.Code
	for _, field := range model.Fields {
		// Comments go first since we like them on top of the type in the generated code
		if field.Comment != "" {
			comment := jen.Comment(field.Comment)
			items = append(items, comment)
		}

		// We detect if a field is Optional or Nullable and construct
		// the correct type for it.
		// Note: Optional goes first since a Nullable field is by definition required
		var idType jen.Code
		if field.Nullable && field.Optional {
			idType = jen.Qual(PackageTypes, "Optional").
				Types(jen.Qual(PackageTypes, "Nullable").
					Types(jen.Qual(field.GoTypePackage, field.GoType)),
				)
		} else if field.Nullable && !field.Optional {
			idType = jen.Qual(PackageTypes, "Nullable").
				Types(jen.Qual(field.GoTypePackage, field.GoType))
		} else if !field.Nullable && field.Optional {
			idType = jen.Qual(PackageTypes, "Optional").
				Types(jen.Qual(field.GoTypePackage, field.GoType))
		} else {
			idType = jen.Qual(field.GoTypePackage, field.GoType)
		}

		id := jen.Id(field.Name).
			Add(idType).
			Tag(field.Annotations)
		items = append(items, id)
	}

	t.Struct(items...)
	modelFile.Add(t)
	return modelFile
}
