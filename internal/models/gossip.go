package models

type Endpoint string

type Gossip struct {
	Message string     `json:message`
	Peers   []Endpoint `json:peers`
}
