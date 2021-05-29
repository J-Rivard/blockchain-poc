package main

import (
	"net/http"
	"os"

	"github.com/J-Rivard/blockchain-poc/internal/api"
	"github.com/J-Rivard/blockchain-poc/internal/config"
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

	config, err := config.New(logger)
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	// publicBlock, _ := pem.Decode([]byte(config.PublicKey))
	// publicKey, err := x509.ParsePKCS1PublicKey(publicBlock.Bytes)

	// privateBlock, _ := pem.Decode([]byte(config.PrivateKey))
	// privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)

	// a, err := pki.Sign2(privateKey, "hello")
	// fmt.Println(a)
	// b := pki.Verify(publicKey, []byte(a), "hello")
	// fmt.Println(b)

	// return

	// en := pki.Encrypt("Hello World", config.PrivateKey)
	// fmt.Println("Encrypted:", en)
	// de := pki.Decrypt(en, config.PublicKey)
	// fmt.Println("Decrypted:", de)

	svc, err := service.New(logger, config.Host, config.InitialPeer)
	if err != nil {
		logger.LogFatal(logging.FormattedLog{
			"action": "startup",
			"error":  err.Error(),
		})
	}

	gossipAPI, err := api.New(svc)
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

		err := http.ListenAndServe(":"+config.HostPort, router)
		if err != nil {
			logger.LogFatal(logging.FormattedLog{
				"action": "startup",
				"error":  err.Error(),
			})
		}
	}()

	logger.LogApplication(logging.FormattedLog{
		"action":   "startup",
		"metadata": "running api on port " + config.HostPort,
	})

	svc.HandleState()

}
