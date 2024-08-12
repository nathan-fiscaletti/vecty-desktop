package container

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/nathan-fiscaletti/lorca"
	"github.com/nathan-fiscaletti/vecty-desktop/container/config"
)

func Main(fileSystem embed.FS) error {
	fs := http.FS(fileSystem)

	// Serve the WebAssembly and support files
	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		fileServer := http.FileServer(fs)
		fileServer.ServeHTTP(w, r)
	})

	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		fileServer := http.FileServer(fs)
		fileServer.ServeHTTP(w, r)
	})

	// Serve the HTML directly from a handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Vecty-Lorca Demo</title>
			<script src="wasm_exec.js"></script>
		</head>
		<body>
			<div id="root"></div>
			<script>
				const go = new Go();
				WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
					go.run(result.instance);
				});
			</script>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	appCfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	appPortString := fmt.Sprintf(":%v", appCfg.Port)
	appUrlString := fmt.Sprintf("http://localhost%v", appPortString)

	// Create a new Lorca UI
	ui, err := lorca.New(appUrlString, "", 800, 600, "--remote-allow-origins=*")
	if err != nil {
		return err
	}
	defer ui.Close()

	// Create a binding for the UI
	err = ui.Bind("getContainerString", func() string {
		return "Hello from the container application!"
	})
	if err != nil {
		return err
	}

	// Start the web server
	go func() {
		log.Fatal(http.ListenAndServe(appPortString, nil))
	}()

	// Wait until UI window is closed
	<-ui.Done()

	return nil
}
