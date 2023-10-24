package server

import (
	"github.com/chanon-mike/ec-microservices/modules/item/itemHandler"
	"github.com/chanon-mike/ec-microservices/modules/item/itemRepository"
	"github.com/chanon-mike/ec-microservices/modules/item/itemUsecase"
)

func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemHandler.NewItemGrpcHandler(itemUsecase)

	_ = itemHttpHandler
	_ = itemGrpcHandler

	item := s.app.Group("/v1/item")

	// Health Check
	item.GET("", s.healthCheckService)
}
