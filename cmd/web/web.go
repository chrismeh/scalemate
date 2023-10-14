package main

import (
	"embed"
	"flag"
	"github.com/chrismeh/scalemate/internal/http"
	"log"
)

//go:embed templates static
var embeddedFiles embed.FS

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "TCP address for the server to listen on")
	flag.Parse()

	app, err := http.NewApplication(*addr, embeddedFiles)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	log.Fatal(err)
}
