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
  
  Once located, copy it to the `./host` directory.

### Running the application

```bash
task run
```

This will compile the WebAssembly code and start the desktop application using `go run`.

### Building the Application

```bash
task build
```

This will compile the WebAssembly code and build the desktop application for the host platform, placing it in `./build`.

## File Structure

- `./app`

Your application code will be placed in this directory. This is where you will write your Vecty code.

- `./host`

This directory is the place in which the Lorca application code is managed. This is separate from the `./app` directory to keep the code separate and to make it easier to manage. It's also used as an intermediate directory when compiling the web assembly code.

- `./build`

When the application is built, the output will be placed in this directory. The output will be the desktop application for the host platform.

