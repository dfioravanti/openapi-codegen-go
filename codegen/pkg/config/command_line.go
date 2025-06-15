package config

import (
	"github.com/dfioravanti/openapi-codegen-go/pkg/models"
)

const (
	OpenapiPathFlagName = "openapi-path"
)

func ParseCommandLine(
	args models.Args,
) (models.Config, error) {

	return models.Config{
		OpenAPIPath: args.OpenApiPath,
	}, nil
}
