package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

		data, err := s.createStateModel()
		if err != nil {
			s.log.LogApplication(logging.FormattedLog{
				"error": err.Error(),
			})
			continue
		}

		resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
		if err != nil {
			s.log.LogApplication(logging.FormattedLog{
				"error": err.Error(),
			})
			continue
		}

		fmt.Println(resp)
	}
}

func (s *Service) createStateModel() ([]byte, error) {
	peerArray := []string{}

	for key, _ := range s.peers {
		peerArray = append(peerArray, string(key))
	}

	return json.Marshal(map[string]interface{}{
		"message": "updating peers",
		"peers":   peerArray,
	})
}
