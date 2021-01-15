// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

//go:generate go run -tags=dev assets_generate.go

func main() {
	err := vfsgen.Generate(
		http.Dir("dist/"),
		vfsgen.Options{
			Filename:     "./pkg/www/assets.go",
			PackageName:  "www",
			VariableName: "ui",
			BuildTags:    "!dev",
		})
	if err != nil {
		log.Fatalln(err)
	}
}
