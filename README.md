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