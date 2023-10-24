package playerHandler

import "github.com/chanon-mike/ec-microservices/modules/player/playerUsecase"

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{playerUsecase}
}
