package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/go-yaml/yaml"
	"github.com/nathan-fiscaletti/lorca"
)

//go:embed config.yaml
var cfgFileSystem embed.FS
var cfg *Config

type Config struct {
	Port int `yaml:"port"`
}

func GetConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	data, err := cfgFileSystem.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	cfg = &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

//go:embed main.wasm wasm_exec.js
var content embed.FS

func main() {
	fs := http.FS(content)

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

	appCfg, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	appPortString := fmt.Sprintf(":%v", appCfg.Port)
	appUrlString := fmt.Sprintf("http://localhost%v", appPortString)

	// Start the web server
	go func() {
		log.Fatal(http.ListenAndServe(appPortString, nil))
	}()

	// Create a new Lorca UI
	ui, err := lorca.New(appUrlString, "", 800, 600, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Wait until UI window is closed
	<-ui.Done()
}
