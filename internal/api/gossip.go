package api

import (
	"fmt"
	"net/http"
)

func (a *GossipAPI) Gossip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gossip")
}
