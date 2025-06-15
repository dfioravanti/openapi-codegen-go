package openapi

import (
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

type PathItems = *orderedmap.Map[string, *v3.PathItem]

func GetPaths(pathItems PathItems) {

	for pathPairs := pathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		path := pathPairs.Value()
		parseOperation(path.Get)
	}
}

func parseOperation(operation *v3.Operation) {
	parseResponses(*operation.Responses)
}

func parseResponses(responses v3.Responses) {
	codes := responses.Codes

	for responsePair := codes.First(); responsePair != nil; responsePair = responsePair.Next() {
		code := responsePair.Key()
		response := responsePair.Value()

		content := response.Content
		for mediaTypePair := content.First(); mediaTypePair != nil; mediaTypePair = mediaTypePair.Next() {
			parseMediaType(*mediaTypePair.Value())
		}
		print(code)
		print(response)
	}

}

func parseMediaType(mediaType v3.MediaType) {
	schema, err := mediaType.Schema.BuildSchema()
	if err != nil {
		panic(err)
	}
	print(schema)
}
