package main

import (
	"net/http"
	"os"

	"github.com/J-Rivard/blockchain-poc/internal/api"
	"github.com/J-Rivard/blockchain-poc/internal/logging"
	"github.com/J-Rivard/blockchain-poc/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func main() {
	logger, err := logging.New(zerolog.ConsoleWriter{Out: os.Stderr}, logging.Debug)
	if err != nil {
		panic(err)
	}

	gossipAPI, err := api.New()
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	svc, err := service.New(&gossipAPI)
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	go func() {
		router := mux.NewRouter()
		router.HandleFunc(api.GossipEndpoint, gossipAPI.Gossip).Methods(http.MethodPost)
		router.HandleFunc(api.SendMoneyEndpoint, gossipAPI.SendMoney).Methods(http.MethodPost)

		err := http.ListenAndServe(":8080", router)
		if err != nil {
			logger.LogFatal(logging.FormattedLog{
				"action": "startup",
				"error":  err.Error(),
			})
		}
	}()

	logger.LogApplication(logging.FormattedLog{
		"action":   "startup",
		"metadata": "running api on port 8080",
	})

	svc.HandleState()

}
