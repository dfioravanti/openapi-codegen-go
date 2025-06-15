package openapi

import (
	"errors"
	"os"

	"github.com/pb33f/libopenapi"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Document = *libopenapi.DocumentModel[v3high.Document]

var (
	CannotCreateDocumentErr  = errors.New("cannot create new libopenapi document")
	CannotParseDocumentV3Err = errors.New("cannot create v3 model from document")
)

func Parse(path string) (Document, error) {
	petstore, _ := os.ReadFile(path)
	document, err := libopenapi.NewDocument(petstore)
	if err != nil {
		return nil, errors.Join(CannotCreateDocumentErr, err)
	}
	docModel, errs := document.BuildV3Model()
	if len(errs) > 0 {
		return nil, errors.Join(CannotParseDocumentV3Err, errors.Join(errs...))
	}

	return docModel, nil
}
