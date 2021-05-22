package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Rivard/blockchain-poc/internal/models"
)

func (a *GossipAPI) Gossip(w http.ResponseWriter, r *http.Request) {
	var gossip models.Gossip

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gossip)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(gossip)
}
