package api

import "github.com/J-Rivard/blockchain-poc/internal/models"

type IService interface {
	Gossip(gossip *models.Gossip) error
	// sendMoney(gossip *models.Gossip) error
}

type GossipAPI struct {
	service IService
}

const (
	GossipEndpoint    = "/gossip"
	SendMoneyEndpoint = "/send-money"
)

func New(svc IService) (GossipAPI, error) {
	return GossipAPI{
		service: svc,
	}, nil
}
