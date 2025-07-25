package events

type EventHandler interface {
	Publish(subject EventSubject, data []byte) error
}
