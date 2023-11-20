package main

import (
	"context"
	"event-service/action"
	"event-service/action/workflow"
	"event-service/event"
	"fmt"
)

func main() {

	// init event manager and listen event
	event.DefaultEM.On(workflow.ReadMq, func(e *event.EventData) error {
		fmt.Printf(" %s event handler: %v\n", workflow.ReadMq, e)
		return nil
	})

	// original data
	data := map[string]string{"subject": "sane's test"}

	// init action manager
	am := action.NewActionManager()
	// register action implements
	am.Register(workflow.Cta, workflow.CtaServiceFactory())
	am.Register(workflow.ReadMq, workflow.ReadMqDataFactory())
	am.Register(workflow.OneId, workflow.OneIdServiceFactory())

	// init action implement and do validate
	actionImpl := am.Get(workflow.ReadMq)
	if err := actionImpl.Validate(context.Background(), data); err != nil {
		panic(err)
	}

	// execute the actual logic
	_, err := actionImpl.Handle(context.Background(), data)
	fmt.Printf(" %s implement result ouput: %v\n", workflow.ReadMq, err)

	// fire event
	err = event.Fire(workflow.ReadMq, data, "sane")
	fmt.Printf(" %s event result ouput: %v\n", workflow.ReadMq, err)

	// Continue processing
	_, _ = am.Get(workflow.Cta).Handle(context.Background(), data)
	_, _ = am.Get(workflow.OneId).Handle(context.Background(), data)
}
