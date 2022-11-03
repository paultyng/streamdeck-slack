package sdk

import (
	"sync"

	"github.com/valyala/fastjson"
)

type ActionHandler func(action, context string, payload *fastjson.Value, deviceId string)

var (
	handlersMu sync.RWMutex

	handlers     map[string][]*eventHandlerInstance
	onceHandlers map[string][]*eventHandlerInstance
)

// EventHandler is an interface for Discord event
type EventHandler interface {
	// Type returns the type of event this handler belongs to.
	Type() string

	// Handle is called whenever an event of Type() happen
	// It is the receivers responsibility to type assert that the interface
	// is the expected struct.
	Handle(interface{})
}

// interfaceEventType is the event handler type for interface{} event
const interfaceEventType = "__INTERFACE__"

// interfaceEventHandler is an event handler for interface{} event
type interfaceEventHandler func(interface{})

// Type returns the event type for interface{} event
func (eh interfaceEventHandler) Type() string {
	return interfaceEventType
}

// Handle is the handler for an interface{} event.
func (eh interfaceEventHandler) Handle(i interface{}) {
	eh(i)
}

// eventHandlerInstance is a wrapper around an event handler, as functions
// cannot be compared directly.
type eventHandlerInstance struct {
	eventHandler EventHandler
}

// addEventHandler adds an event handler that will be fired anytime
// the Discord WSAPI matching eventHandler.Type() fire
func addEventHandler(eventHandler EventHandler) func() {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	if handlers == nil {
		handlers = map[string][]*eventHandlerInstance{}
	}

	ehi := &eventHandlerInstance{eventHandler}
	handlers[eventHandler.Type()] = append(handlers[eventHandler.Type()], ehi)

	return func() {
		removeEventHandlerInstance(eventHandler.Type(), ehi)
	}
}

// addEventHandler adds an event handler that will be fired the next time
// the Stream Deck matching eventHandler.Type() fire
func addEventHandlerOnce(eventHandler EventHandler) func() {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	if onceHandlers == nil {
		onceHandlers = map[string][]*eventHandlerInstance{}
	}

	ehi := &eventHandlerInstance{eventHandler}
	onceHandlers[eventHandler.Type()] = append(onceHandlers[eventHandler.Type()], ehi)

	return func() {
		removeEventHandlerInstance(eventHandler.Type(), ehi)
	}
}

// AddHandler allows you to add an event handler that will be fired anytime
// the Stream Deck event the function matches is received
// The second parameter is a pointer to a struct
// corresponding to the event for which you want to listen.
//
// eg:
//
//	AddHandler(func(e *sdk.KeyDownEvent) {
//	})
//
// or:
//
//	AddHandler(func(e *sdk.SendToPluginEvent) {
//	})
//
// The return value of this method is a function, that when called will remove the
// event handler.
func AddHandler(handler interface{}) func() {
	eh := handlerForInterface(handler)

	if eh == nil {
		return func() {}
	}

	return addEventHandler(eh)
}

// AddHandlerOnce allows you to add an event handler that will be fired the next time
// the Discord WSAPI event that matches the function fire
// See AddHandler for more detail
func AddHandlerOnce(handler interface{}) func() {
	eh := handlerForInterface(handler)

	if eh == nil {
		return func() {}
	}

	return addEventHandlerOnce(eh)
}

// removeEventHandler instance removes an event handler instance.
func removeEventHandlerInstance(t string, ehi *eventHandlerInstance) {
	handlersMu.Lock()
	defer handlersMu.Unlock()

	rhandlers := handlers[t]
	for i := range rhandlers {
		if rhandlers[i] == ehi {
			handlers[t] = append(rhandlers[:i], rhandlers[i+1:]...)
		}
	}

	ronceHandlers := onceHandlers[t]
	for i := range ronceHandlers {
		if ronceHandlers[i] == ehi {
			onceHandlers[t] = append(ronceHandlers[:i], rhandlers[i+1:]...)
		}
	}
}

// Handles calling permanent and once handlers for an event type.
func handle(t string, i interface{}) {
	for _, eh := range handlers[t] {
		eh.eventHandler.Handle(i)
	}

	if len(onceHandlers[t]) > 0 {
		for _, eh := range onceHandlers[t] {
			eh.eventHandler.Handle(i)
		}
		onceHandlers[t] = nil
	}
}

// Handles an event type by calling internal methods, firing handlers and firing the
// interface{} event.
func handleEvent(t string, i interface{}) {
	handlersMu.RLock()
	defer handlersMu.RUnlock()

	// Then they are dispatched to anyone handling interface{} event
	handle(interfaceEventType, i)

	// Finally they are dispatched to any typed handler
	handle(t, i)
}
