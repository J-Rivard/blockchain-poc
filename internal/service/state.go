package service

import (
	"fmt"
	"time"
)

func (s *Service) HandleState() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
