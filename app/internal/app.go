//go:build js && wasm
// +build js,wasm

package internal

import (
	"github.com/hexops/vecty"
	"github.com/nathan-fiscaletti/vecty-desktop/app/internal/components"
)

func Main() error {
	// Create and render the Vecty component
	vecty.RenderBody(&components.ExampleComponent{})

	return nil
}
