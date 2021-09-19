package http

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	server     *http.Server
	infoLog    *log.Logger
	errorLog   *log.Logger
	templateFS fs.FS
}

func NewApplication(embeddedFiles embed.FS) (Application, error) {
	app := Application{
		server: &http.Server{
			Addr:         ":8080",
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	templateFS, err := fs.Sub(embeddedFiles, "templates")
	if err != nil {
		return Application{}, err
	}
	app.templateFS = templateFS

	router := http.NewServeMux()
	router.HandleFunc("/scale", app.handleGetScale)
	router.HandleFunc("/", app.handleGetIndex)
	app.server.Handler = router

	return app, nil
}

func (a Application) Run() error {
	a.infoLog.Printf("starting application at port %s", a.server.Addr)
	return a.server.ListenAndServe()
}
