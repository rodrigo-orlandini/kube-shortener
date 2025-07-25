package stubMessaging

import "rodrigoorlandini/urlshortener/shortener/internal/domain/events"

type StubEventHandler struct {
	Events map[events.EventSubject][][]byte
}

func NewStubEventHandler() events.EventHandler {
	return &StubEventHandler{
		Events: make(map[events.EventSubject][][]byte),
	}
}

func (e *StubEventHandler) Publish(subject events.EventSubject, data []byte) error {
	e.Events[subject] = append(e.Events[subject], data)

	return nil
}
