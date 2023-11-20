package workflow

import (
	"context"
	"event-service/action"
)

const ReadMq string = "read-mq-data"

type ReadMqData struct {
}

func ReadMqDataFactory() action.ActionFactory {
	return func() action.ActionInterface {
		return &ReadMqData{}
	}
}

func (r *ReadMqData) Validate(ctx context.Context, data any) error {
	return nil
}

func (r *ReadMqData) Handle(ctx context.Context, data any) (any, error) {
	return nil, nil
}
