package server

import (
	"github.com/chanon-mike/ec-microservices/modules/auth/authHandler"
	"github.com/chanon-mike/ec-microservices/modules/auth/authRepository"
	"github.com/chanon-mike/ec-microservices/modules/auth/authUsecase"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authHandler.NewAuthGrpcHandler(authUsecase)

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/v1/auth")

	// Health Check
	auth.GET("", s.healthCheckService)
}
