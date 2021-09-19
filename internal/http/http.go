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

	staticFiles, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		return Application{}, err
	}

	router := http.NewServeMux()
	router.HandleFunc("/", app.handleGetIndex)
	router.HandleFunc("/scale", app.handleGetScale)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
	app.server.Handler = router

	return app, nil
}

func (a Application) Run() error {
	a.infoLog.Printf("starting application at port %s", a.server.Addr)
	return a.server.ListenAndServe()
}

func (a Application) internalServerError(err error, w http.ResponseWriter) {
	a.errorLog.Println(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (a Application) badRequest(err error, w http.ResponseWriter) {
	a.errorLog.Println(err)
	http.Error(w, "Bad Request", http.StatusBadRequest)
}
