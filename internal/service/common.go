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
	peers := make(map[models.Endpoint]struct{})

	if initPeer != "" {
		peers[models.Endpoint(initPeer)] = struct{}{}
	}

	return &Service{
		host:  models.Endpoint(host),
		peers: peers,
		log:   log,
	}, nil
}
