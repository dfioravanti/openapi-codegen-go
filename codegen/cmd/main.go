package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	. "github.com/dave/jennifer/jen"
	"github.com/dfioravanti/openapi-codegen-go/pkg/config"
	"github.com/dfioravanti/openapi-codegen-go/pkg/models"
	"github.com/dfioravanti/openapi-codegen-go/pkg/openapi"
)

const (
	typeImportString = "github.com/dfioravanti/openapi-codegen-go/types"
)

func main() {

	var args models.Args
	arg.MustParse(&args)

	cfg, err := config.ParseCommandLine(args)
	if err != nil {
		fmt.Printf("cannot parse command line configs: %s", err)
		return
	}

	openAPIDocument, err := openapi.Parse(cfg.OpenAPIPath)
	if err != nil {
		fmt.Printf("cannot parse openapi document: %s", err)
		return
	}

	pathItems := openAPIDocument.Model.Paths.PathItems
	openapi.GetPaths(pathItems)

	f := NewFile("a")

	t := Type().Id("foo").Struct(
		Id("bar").Qual(typeImportString, "Optional").Types(Qual(typeImportString, "Nullable").Types(String())),
	)
	f.Add(t)
	fmt.Printf("%#v", f)
}
