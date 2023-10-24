package server

import (
	"github.com/chanon-mike/ec-microservices/modules/payment/paymentHandler"
	"github.com/chanon-mike/ec-microservices/modules/payment/paymentRepository"
	"github.com/chanon-mike/ec-microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	paymentRepository := paymentRepository.NewPaymentRepository(s.db)
	paymentUsecase := paymentUsecase.NewPaymentUsecase(paymentRepository)
	paymentHttpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, paymentUsecase)
	paymentQueueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, paymentUsecase)

	_ = paymentHttpHandler
	_ = paymentQueueHandler

	payment := s.app.Group("/v1/payment")

	// Health Check
	payment.GET("", s.healthCheckService)
}
