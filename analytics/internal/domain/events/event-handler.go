package events

type EventHandler interface {
	Subscribe(subject EventSubject, handler func(data []byte) error) error
}
