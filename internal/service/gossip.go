package service

import (
	"github.com/J-Rivard/blockchain-poc/internal/logging"
	"github.com/J-Rivard/blockchain-poc/internal/models"
)

func (s *Service) Gossip(gossip *models.Gossip) error {
	s.log.LogApplication(logging.FormattedLog{
		"gossip": gossip.Message,
	})

	s.updatePeers(gossip.Peers)

	return nil
}

func (s *Service) updatePeers(peers []models.Endpoint) {
	for _, peer := range peers {
		if string(peer) == string(s.host) {
			continue
		}
		s.peers[peer] = struct{}{}
	}

	peerString := ""
	for key, _ := range s.peers {
		peerString += string(key) + ","
	}

	s.log.LogApplication(logging.FormattedLog{
		"action": "updatedPeers",
		"peers":  peerString,
	})
}
