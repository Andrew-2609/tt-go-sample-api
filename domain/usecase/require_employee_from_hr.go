package usecase

import (
	"context"
	"time"
	"tt-go-sample-api/domain/usecase/dto"
	"tt-go-sample-api/external/messaging"
)

// RequireEmployeeFromHRUseCase is a concrete implementation
// of RequireEmployeeFromHRUseCaseInterface. It shall be
// used to require new employees from the HR (Human Resources).
type RequireEmployeeFromHRUseCase struct {
	messageProducer messaging.MessageProducer
}

// NewRequireEmployeeFromHRUseCase returns a pointer of
// RequireEmployeeFromHRUseCase, with the given message producer.
func NewRequireEmployeeFromHRUseCase(messageProducer messaging.MessageProducer) *RequireEmployeeFromHRUseCase {
	return &RequireEmployeeFromHRUseCase{messageProducer: messageProducer}
}

// Execute tries to require an employee from the HR.
func (uc *RequireEmployeeFromHRUseCase) Execute(ctx context.Context, input dto.RequireEmployeeFromHRInputDTO) (dto.RequireEmployeeFromHROutputDTO, error) {
	err := uc.messageProducer.Produce(ctx, messaging.Message{
		ID:      time.Now().Format(time.DateOnly),
		GroupID: input.Stack,
		Payload: input,
	})

	if err != nil {
		return dto.RequireEmployeeFromHROutputDTO{}, err
	}

	return dto.RequireEmployeeFromHROutputDTO{
		Message: "Employee successfully requested! See the employees list in about 5 seconds!",
	}, nil
}
