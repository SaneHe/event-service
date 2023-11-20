package workflow

import (
	"context"
	"event-service/action"
	"log"
	"time"
)

const ReadMq string = "read-mq-data"

type (
	ReadMqData struct {
	}
	
	ReadMqEvent struct {
		Time time.Time
		data string
	}
)

func (r *ReadMqData) Validate(ctx context.Context, data any) error {
	return nil
}

func (r *ReadMqData) Handle(ctx context.Context, data any) error {
	return nil
}

func ReadMqDataFactory() action.ActionFactory {
	return func() action.ActionInterface {
		return &ReadMqData{}
	}
}

func (e *ReadMqEvent) Handle(ctx context.Context) {
	log.Printf("read mq data: %+v\n", e)
}
