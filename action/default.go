package action

import (
	"context"
)

type (
	ActionInterface interface {
		Validate(ctx context.Context, data any) error
		Handle(ctx context.Context, data any) error
	}
	
	ActionFactory func() ActionInterface
	
	ActionManager struct {
		actions map[string]ActionInterface
	}
)

func (am *ActionManager) Register(name string, factory ActionFactory) {
	am.actions[name] = factory()
}

func (am *ActionManager) Get(name string) ActionInterface {
	return am.actions[name]
}

func NewActionManager() *ActionManager {
	return &ActionManager{
		actions: make(map[string]ActionInterface),
	}
}
