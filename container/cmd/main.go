package main

import (
	"embed"

	container "github.com/nathan-fiscaletti/vecty-desktop/container/internal"
)

//go:embed *.wasm *.html *.css *.js
var content embed.FS

func main() {
	if err := container.Main(content); err != nil {
		panic(err)
	}
}
