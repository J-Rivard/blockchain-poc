package service

import (
	"github.com/J-Rivard/blockchain-poc/internal/logging"
	"github.com/J-Rivard/blockchain-poc/internal/models"
)

type Service struct {
	host  models.Endpoint
	peers map[models.Endpoint]struct{}
	log   *logging.Log
}

func New(log *logging.Log, host string, initPeer string) (*Service, error) {
	return &Service{
		host: models.Endpoint(host),
		peers: map[models.Endpoint]struct{}{
			models.Endpoint(initPeer): {},
		},
		log: log,
	}, nil
}
