package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SureshKumar-U/Contacts_management/Routes"
)

type appConfig struct {
	AppName string
	Server  server
	Debug   bool
	ErrLog  *log.Logger
	InfoLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func NewAppConig() *appConfig {
	server := server{
		host: "localhost",
		port: "8000",
		url:  "http://localhost:8000",
	}
	app := &appConfig{
		AppName: "contact Management",
		Server:  server,
		Debug:   true,
		ErrLog:  log.New(os.Stdout, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile),
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
	}
	return app

}

func (app *appConfig) Run() error {
	err := app.ListenAndServe()
	return err

}

func (a *appConfig) Route() http.Handler {
	mux := http.NewServeMux()
	Routes.ContactRoutes(mux)
	Routes.UserRoutes(mux)
	return mux
}

func (a *appConfig) ListenAndServe() error {

	host := fmt.Sprintf("%s:%s", a.Server.host, a.Server.port)
	server := http.Server{
		Addr:        host,
		Handler:     a.Route(),
		ReadTimeout: 300 * time.Millisecond,
	}
	a.InfoLog.Printf("server listening on %s", a.Server.port)
	return server.ListenAndServe()

}
