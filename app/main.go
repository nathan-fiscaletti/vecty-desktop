//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/hexops/vecty"

	"github.com/nathan-fiscaletti/vecty-meets-lorca/app/components"
)

func main() {
	// Create and render the Vecty component
	vecty.RenderBody(&components.ExampleComponent{})
}
