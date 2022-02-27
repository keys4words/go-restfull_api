package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerFields() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterFields() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey! This is REST API..."))
	})
}
