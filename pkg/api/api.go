package api

import (
	"fmt"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/golang-template/pkg/types"
	"net/http"
	"os"
	"time"

	"github.com/golang-template/pkg/loggers/logiface"
)

const version = "version-1.0\n"

type Api struct {
	router  *chi.Mux
	logger  logiface.Logiface
	host    types.HttpHost
	port    types.ConfHttpPort
	timeout time.Duration
}

func (a *Api) StartRouter() {
	err := a.host.Validate()
	if err != nil {
		a.logger.Error(fmt.Sprintf("Wrong host address string. Must be <host>:<port> . Error is [%v]", err))
		os.Exit(1)
	}

	hostAddress := fmt.Sprintf("%s:%s", a.host.String(), a.port.String())
	a.logger.Info(fmt.Sprintf("Starting chi router on host [%s] and  port [%s]", a.host.String(), a.port.String()))
	a.router = chi.NewRouter()

	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	//a.router.Use(middleware.Logger)
	a.router.Use(a.logger.ServeHTTP) //Custom logger from interface
	a.router.Use(middleware.Recoverer)
	a.router.Use(render.SetContentType(render.ContentTypeJSON))

	a.router.Use(middleware.Timeout(a.timeout * time.Second))

	a.router.Get("/version", getVersion)

	err = http.ListenAndServe(hostAddress, a.router)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Unable to start router. [%v]", err))
	}
}

func NewApi(logger logiface.Logiface, host types.HttpHost, port types.ConfHttpPort, timeout time.Duration) *Api {
	api := Api{
		logger:  logger,
		host:    host,
		port:    port,
		timeout: timeout,
	}
	return &api
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
}
