package main

import (
	"embed"
	"github.com/chrismeh/scalemate/internal/http"
	"log"
)

//go:embed templates static
var embeddedFiles embed.FS

func main() {
	app, err := http.NewApplication(embeddedFiles)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	log.Fatal(err)
}
