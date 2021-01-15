package main

import (
	"fmt"
	"net/http"

	"github.com/webdevliteral/FractalGeneration/transport"
)

func main() {
	http.HandleFunc("/fractal", transport.FractalImageHandler)

	fmt.Println("starting http server")

	if err := http.ListenAndServe(":40000", nil); err != nil {
		fmt.Printf("unable to listen on port 40000: %v\n", err)
		return
	}
}
