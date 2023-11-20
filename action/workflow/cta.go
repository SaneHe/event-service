package workflow

import (
	"context"
	"event-service/action"
)

const Cta string = "cta-service"

type CtaService struct {
}

func CtaServiceFactory() action.ActionFactory {
	return func() action.ActionInterface {
		return &CtaService{}
	}
}

func (c *CtaService) Validate(ctx context.Context, data any) error {
	return nil
}

func (c *CtaService) Handle(ctx context.Context, data any) (any, error) {
	return nil, nil
}
