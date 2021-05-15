package api

type GossipAPI struct {
}

const (
	GossipEndpoint    = "/gossip"
	SendMoneyEndpoint = "/send-money"
)

func New() (GossipAPI, error) {
	return GossipAPI{}, nil
}
