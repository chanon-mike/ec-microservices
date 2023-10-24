package paymentUsecase

import "github.com/chanon-mike/ec-microservices/modules/payment/paymentRepository"

type (
	PaymentUsecaseService interface{}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{paymentRepository}
}
