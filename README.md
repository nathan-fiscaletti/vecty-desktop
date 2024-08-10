# Vecty meets Lorca

This is a simple example of how to use [Vecty](https://github.com/gopherjs/vecty) with [Lorca](https://github.com/zserge/lorca) to create a desktop application.

## How to run

### Preqrequisites

- **TaskFile**

  Task is used for running the build and run commands.
  [Installation Instructions](https://taskfile.dev/installation/)

- **`wasm_exec.js`**

  You need to grab the `wasm_exec.js` file from your `GOROOT` in order to properly build your application. You can find it in the `misc/wasm` directory.

  If you do not know where your `GOROOT` is, you can run the following command:

  ```bash
  go env GOROOT
  ```
  
  Once located, copy it to the `./build` directory.

### Running the application

```bash
task run
```

This will compile the WebAssembly code and start the desktop application.