package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var tfs http.FileSystem = http.Dir("templates")
	err := vfsgen.Generate(tfs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}
}
