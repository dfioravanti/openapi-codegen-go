package models

type Args struct {
	OpenApiPath string `arg:"--openapi-path,-i,required" help:"the openapi input file path"`
}
