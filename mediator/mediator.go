package main

type CallbackFn func(data any)

type Event struct {
	event    string
	callback CallbackFn
}

type Mediator struct {
	handlers []*Event
}

func (m *Mediator) Register(event string, callback CallbackFn) {
	m.handlers = append(m.handlers, &Event{event: event, callback: callback})
}

func (m *Mediator) Notify(event string, data any) {
	for _, h := range m.handlers {
		if h.event == event {
			h.callback(data)
		}
	}
}
