package models

type Field struct {
	Name    string
	Comment string
	// Important: we use this models.ModelsPackageName to generate built in names
	// since jennifer does not render this name as part of a type,
	// so Qual(ModelsPackageName, "int") becomes int.
	GoTypePackage string
	GoType        string
	// Optional marks if the attribute is optional and we will prepend the attribute
	// with Optional[Type].
	// If it is NOT optional then we expect a required annotation in the annotations
	// so that it can be properly validated when we parse the body.
	Optional bool
	// Nullable marks if the attribute is nullable and we will prepend the attribute
	// with Nullable[Type].
	// A attribute field is required if it is not explicitly marked as optional.
	Nullable    bool
	Annotations map[string]string
}

type Model struct {
	Name   string
	Fields []Field
}
