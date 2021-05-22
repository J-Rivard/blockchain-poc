package service

import "net/http"

type IGossipAPI interface {
	SendMoney(w http.ResponseWriter, r *http.Request)
	Gossip(w http.ResponseWriter, r *http.Request)
}

type Endpoint string

type Service struct {
	gossipAPI IGossipAPI
	host      Endpoint
	peers     []Endpoint
}

func New(gossipAPI IGossipAPI, host string, initPeer string) (*Service, error) {
	return &Service{
		host:  Endpoint(host),
		peers: []Endpoint{Endpoint(initPeer)},

		gossipAPI: gossipAPI,
	}, nil
}
