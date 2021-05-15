package service

import "net/http"

type IGossipAPI interface {
	SendMoney(w http.ResponseWriter, r *http.Request)
	Gossip(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	gossipAPI IGossipAPI
}

func New(gossipAPI IGossipAPI) (*Service, error) {
	return &Service{
		gossipAPI: gossipAPI,
	}, nil
}
