//go:build js
// +build js

package main

import "github.com/nathan-fiscaletti/vecty-desktop/app/internal"

func main() {
	if err := internal.Main(); err != nil {
		panic(err)
	}
}
