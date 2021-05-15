package api

import (
	"fmt"
	"net/http"
)

func (a *GossipAPI) SendMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Println("send money")
}
