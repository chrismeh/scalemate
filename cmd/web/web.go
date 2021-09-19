package main

import (
	"github.com/chrismeh/scalemate/internal/http"
	"log"
)

func main() {
	app := http.NewApplication()
	log.Fatal(app.Run())
}
