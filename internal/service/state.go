package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/J-Rivard/blockchain-poc/internal/logging"
)

func (s *Service) HandleState() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			s.updateState()
		}
	}
}

func (s *Service) updateState() {
	for peer, _ := range s.peers {
		endpoint := string(peer) + "/gossip"

		err := s.gossipPeers(endpoint)
		if err != nil {
			s.log.LogApplication(logging.FormattedLog{
				"error": err.Error() + ", removing from peer network",
			})

			delete(s.peers, peer)
		}
	}
}

func (s *Service) gossipPeers(endpoint string) error {
	data, err := s.createStateModel()
	if err != nil {
		return err
	}

	urlEndpoint, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	resp, err := http.Post(urlEndpoint.String(), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}

func (s *Service) createStateModel() ([]byte, error) {
	peerArray := []string{string(s.host)}

	for key, _ := range s.peers {
		peerArray = append(peerArray, string(key))
	}

	return json.Marshal(map[string]interface{}{
		"message": "updating peers",
		"peers":   peerArray,
	})
}
