package event

import (
	"bytes"
	"strings"
	"sync"
)

// Wildcard event name
const Wildcard = "*"

type (
	// HandlerFunc event handler func define
	HandlerFunc func(e *EventData) error

	// Manager struct
	Manager struct {
		pool  sync.Pool
		names map[string]int
		// storage event handlers
		handlers map[string][]HandlerFunc
	}
)

// NewManager create Manager instance
func NewManager() *Manager {
	em := &Manager{
		names:    make(map[string]int),
		handlers: make(map[string][]HandlerFunc),
	}

	// set pool creator
	em.pool.New = func() any {
		return &EventData{}
	}

	return em
}

// On register a event handler
func (em *Manager) On(name string, handler HandlerFunc) {
	if name = strings.TrimSpace(name); name == "" {
		panic("event name cannot be empty")
	}

	if ls, ok := em.handlers[name]; ok {
		em.names[name]++
		em.handlers[name] = append(ls, handler)
	} else { // first add.
		em.names[name] = 1
		em.handlers[name] = []HandlerFunc{handler}
	}
}

// MustFire fire handlers by name. will panic on error
func (em *Manager) MustFire(name string, args ...any) {
	err := em.Fire(name, args...)
	if err != nil {
		panic(err)
	}
}

// Fire handlers by name
func (em *Manager) Fire(name string, args ...any) (err error) {
	handlers, ok := em.handlers[name]
	if !ok {
		return
	}

	e := em.pool.Get().(*EventData)
	e.init(name, args)

	// call event handlers
	err = em.doFire(e, handlers)

	e.reset()
	em.pool.Put(e)
	return
}

func (em *Manager) doFire(e *EventData, handlers []HandlerFunc) (err error) {
	err = em.callHandlers(e, handlers)
	if err != nil || e.IsAborted() {
		return
	}

	// group listen "app.*"
	// groupName :=
	// Wildcard event handler
	if em.HasEvent(Wildcard) {
		err = em.callHandlers(e, em.handlers[Wildcard])
	}

	return
}

func (em *Manager) callHandlers(e *EventData, handlers []HandlerFunc) (err error) {
	for _, handler := range handlers {
		err = handler(e)
		if err != nil || e.IsAborted() {
			return
		}
	}
	return
}

// HasEvent has event check
func (em *Manager) HasEvent(name string) bool {
	_, ok := em.names[name]
	return ok
}

// GetEventHandlers get handlers and handlers by name
func (em *Manager) GetEventHandlers(name string) (es []HandlerFunc) {
	es, _ = em.handlers[name]
	return
}

// EventHandlers get all event handlers
func (em *Manager) EventHandlers() map[string][]HandlerFunc {
	return em.handlers
}

// EventNames get all event names
func (em *Manager) EventNames() map[string]int {
	return em.names
}

// String convert to string.
func (em *Manager) String() string {
	buf := new(bytes.Buffer)
	for name, hs := range em.handlers {
		buf.WriteString(name + " handlers:\n ")
		for _, h := range hs {
			buf.WriteString(funcName(h))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

// ClearHandlers clear handlers by name
func (em *Manager) ClearHandlers(name string) bool {
	_, ok := em.names[name]
	if ok {
		delete(em.names, name)
		delete(em.handlers, name)
	}
	return ok
}

// Clear all handlers' info.
func (em *Manager) Clear() {
	em.names = map[string]int{}
	em.handlers = map[string][]HandlerFunc{}
}
