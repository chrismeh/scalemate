package http

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	server   *http.Server
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewApplication() Application {
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

	router := http.NewServeMux()
	router.HandleFunc("/scale", app.handleGetScale)
	app.server.Handler = router

	return app
}

func (a Application) Run() error {
	a.infoLog.Printf("starting application at port %s", a.server.Addr)
	return a.server.ListenAndServe()
}
