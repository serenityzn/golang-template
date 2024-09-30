package api

import (
	"fmt"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/golang-template/pkg/loggers/logiface"
)

const version = "version-1.0\n"

type Api struct {
	router *chi.Mux
	logger logiface.Logiface
}

func (a *Api) StartRouter(hostAddress string, timeout time.Duration) {
	a.logger.Info("Starting chi router for rest api.")
	if !ValidateHostAddress(hostAddress) {
		a.logger.Error("Wrong host address string. Must be <host>:<port> .")
		os.Exit(1)
	}
	a.router = chi.NewRouter()

	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	//a.router.Use(middleware.Logger)
	a.router.Use(a.logger.ServeHTTP)
	a.router.Use(middleware.Recoverer)

	a.router.Use(middleware.Timeout(timeout * time.Second))

	a.router.Get("/version", getVersion)

	err := http.ListenAndServe(hostAddress, a.router)
	if err != nil {
		a.logger.Error(fmt.Sprintf("Unable to start router. [%v]", err))
	}
}

func NewApi(logger logiface.Logiface) *Api {
	api := Api{
		logger: logger,
	}
	return &api
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
}

func ValidateHostAddress(address string) bool {
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9]):([0-9]{1,5})$`)
	return re.MatchString(address)
}
