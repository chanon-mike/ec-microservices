package server

import (
	"github.com/chanon-mike/ec-microservices/modules/player/playerHandler"
	"github.com/chanon-mike/ec-microservices/modules/player/playerRepository"
	"github.com/chanon-mike/ec-microservices/modules/player/playerUsecase"
)

func (s *server) playerService() {
	playerRepository := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(playerUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(playerUsecase)
	playerQueueHandler := playerHandler.NewPlayerQueueHandler(playerUsecase)

	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = playerQueueHandler

	player := s.app.Group("/v1/player")

	// Health Check
	player.GET("", s.healthCheckService)
}
