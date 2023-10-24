package server

import (
	"github.com/chanon-mike/ec-microservices/modules/inventory/inventoryHandler"
	"github.com/chanon-mike/ec-microservices/modules/inventory/inventoryRepository"
	"github.com/chanon-mike/ec-microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryUsecase)
	inventoryQueueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	_ = inventoryHttpHandler
	_ = inventoryGrpcHandler
	_ = inventoryQueueHandler

	inventory := s.app.Group("/v1/inventory")

	// Health Check
	inventory.GET("", s.healthCheckService)
}
