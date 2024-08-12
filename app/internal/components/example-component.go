package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// ExampleComponent is a simple Vecty component.
type ExampleComponent struct {
	vecty.Core

	text string
}

func (c *ExampleComponent) onClick(event *vecty.Event) {
	c.text = "Hello, world!"
	vecty.Rerender(c)
}

func (c *ExampleComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Button(
			vecty.Text("Click me!"),
			vecty.Markup(
				event.Click(c.onClick),
			),
		),
		vecty.Text(c.text),
	)
}
