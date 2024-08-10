package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// ExampleComponent is a simple Vecty component.
type ExampleComponent struct {
	vecty.Core
}

func (c *ExampleComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello from Vecty in WebAssembly!"),
	)
}
