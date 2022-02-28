package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keys4words/go-restfull_api/standardWebserver/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerFields(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	api.configureRouterFields()
	if err := api.configureStorageFields(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
