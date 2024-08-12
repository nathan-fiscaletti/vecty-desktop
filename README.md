# Vecty Desktop

This template enables you to create a desktop application using [Vecty](https://github.com/gopherjs/vecty) and manage it entirely in Go. By leveraging [Lorca](https://github.com/zserge/lorca), this setup handles both the host application and the WebAssembly code, making it easy to build desktop apps with Go and WebAssembly.

## Getting Started

### Prerequisites

- **TaskFile**

  To manage build and run commands, this template uses Task. You can find installation instructions [here](https://taskfile.dev/installation/).

- **`wasm_exec.js`**

  You’ll need the `wasm_exec.js` file from your `GOROOT` to properly build your application. This file is located in the `misc/wasm` directory of your Go installation.

  If you’re unsure of your `GOROOT`, run:

  ```bash
  go env GOROOT
  ```

  Once you have located the `wasm_exec.js` file, copy it to the `./container/cmd` directory of this template. 
  
  >This file is only required for building the container application and is not needed for running the application.

### Running the Application

To compile the WebAssembly code and start the desktop application, use:

```bash
task run
```

### Building the Application

To build the WebAssembly code and package the desktop application for the host platform, use:

```bash
task build
```

The build output will be placed in the `./dist` directory.

## Project Structure

- **`./app`**

  This directory contains the code used to build the WebAssembly application. This is where you will make changes to the WebAssembly application.

- **`./container`**

  This directory contains the code used to build the desktop application. This is where you will make changes to the desktop application.

- **`./dist`**

  This directory stores the final build output, containing the packaged desktop application for the host platform.

## Inter Process Communication

To expose a function from the container application to the WebAssembly application, use the `bind` function on the Lorca `ui` object. This function takes the name of the function and the function itself as arguments.

```go
// Create a binding for the UI
err := ui.Bind("getContainerString", func() string {
  return "Hello from the container application!"
})
if err != nil {
  log.Fatal(err)
}
```

In order to call this function from within your WebAssembly code, use the `syscall/js` package.

```go
import "syscall/js"
```

You can then call the function using the `js.Global().Call` function.

```go
containerStringPromise := js.Global().Call("getContainerString")
```

Keep in mind that when using `ui.Bind` in Lorca, the function will be bound as an asynchronous function. This means that when calling the function from the WebAssembly code, you will receive a Promise object that you can use to handle the result.

```go
containerStringPromise := js.Global().Call("getContainerString")
containerStringPromise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
  // Extract the result from the Promise
  res := args[0].String()

  // Update the UI with the result
  c.text = res
  vecty.Rerender(c)
  return nil
}))
```

There is also a utility function built into this package to make calling these functions easier. You can use the `util.CallPromise()` function to call the function and handle the result in one step.

```go
util.CallPromise("getContainerString", util.PromiseHandler{
  Resolve: func(value js.Value) {
    c.text = value.String()
  },
  Reject: func(err js.Error) {
    c.text = err.Error()
  },
  Finally: func() {
    vecty.Rerender(c)
  },
})
```

You can also call Javascript functions in your WebAssembly code from the container application using the `ui.Eval` function.

```go
ui.Eval(`console.log("Hello from the container!")`)
```