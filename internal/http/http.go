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
	router := http.NewServeMux()
	srv := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	return Application{
		server:   srv,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (a Application) Run() error {
	a.infoLog.Printf("starting application at port %s", a.server.Addr)
	return a.server.ListenAndServe()
}
