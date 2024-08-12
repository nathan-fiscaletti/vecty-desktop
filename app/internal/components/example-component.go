package components

import (
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/nathan-fiscaletti/vecty-desktop/app/internal/util"
)

// ExampleComponent is a simple Vecty component.
type ExampleComponent struct {
	vecty.Core

	text string
}

func (c *ExampleComponent) onClick(event *vecty.Event) {
	// Simulate calling
	// getContainerString().then(value => ...)
	util.CallPromise("getContainerString", util.PromiseHandler{
		Resolve: func(values []js.Value) {
			c.text = values[0].String()
		},
		Reject: func(err js.Error) {
			c.text = err.Error()
		},
		Finally: func() {
			vecty.Rerender(c)
		},
	})
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
