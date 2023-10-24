package authUsecase

import "github.com/chanon-mike/ec-microservices/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository}
}
