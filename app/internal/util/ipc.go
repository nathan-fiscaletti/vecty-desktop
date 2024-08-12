package util

import (
	"fmt"
	"syscall/js"
)

type PromiseHandler struct {
	Resolve func([]js.Value)
	Reject  func(js.Error)
	Finally func()
}

// CallHostFunction is a utility function that allows you to call a function
// that is defined in the host environment (the Go code running in the container).
func CallPromise(name string, handler PromiseHandler, args ...any) {
	// Wrap the handlers in js.Func and release them after use
	thenFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler.Resolve(args)
		return nil
	})

	catchFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler.Reject(js.Error{Value: args[0]})
		return nil
	})

	var finallyFunc js.Func
	finallyFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if handler.Finally != nil {
			handler.Finally()
		}

		thenFunc.Release()
		catchFunc.Release()
		finallyFunc.Release()
		return nil
	})

	promise := js.Global().Call(name, args...)
	if promise.Type() != js.TypeObject {
		handler.Reject(js.Error{Value: js.ValueOf(fmt.Errorf("expected a promise, got %v", promise.Type()))})
		return
	}

	promise.Call("then", thenFunc)
	promise.Call("catch", catchFunc)
	promise.Call("finally", finallyFunc)
}
