package app

import (
	"net/http"
	"time"

	"github.com/braddle/go-http-template/accesslog"
	"github.com/braddle/go-http-template/clock"

	"github.com/braddle/go-http-template/rest"
	"github.com/gorilla/mux"
)

type App struct {
	r *mux.Router
}

func New(r *mux.Router) *App {
	a := &App{r: r}
	a.init()

	return a
}

func (a *App) init() {
	al := accesslog.New(clock.New())
	a.r.Use(al.Logger)

	a.r.Handle("/healthcheck", a.getHealthCheckHandle()).Methods(http.MethodGet)
	a.r.NotFoundHandler = a.r.NewRoute().Handler(a.getNotFoundHandle()).GetHandler()
}

func (a *App) Run(host string) error {
	srv := http.Server{
		Addr:         host,
		Handler:      a.r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (a *App) getHealthCheckHandle() http.Handler {
	return rest.HealthCheck{}
}

func (a *App) getNotFoundHandle() http.Handler {
	return rest.NotFound{}
}
