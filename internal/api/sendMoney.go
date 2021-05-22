package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *GossipAPI) SendMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Println("send money")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", reqBody)
}
