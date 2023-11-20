package workflow

import (
	"context"
	"event-service/action"
)

const OneId string = "one-id-service"

type OneIdService struct {
}

func OneIdServiceFactory() action.ActionFactory {
	return func() action.ActionInterface {
		return &OneIdService{}
	}
}

func (o *OneIdService) Validate(ctx context.Context, data any) error {
	return nil
}

func (o *OneIdService) Handle(ctx context.Context, data any) (any, error) {
	return nil, nil
}
