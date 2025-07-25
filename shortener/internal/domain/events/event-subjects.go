package events

type EventSubject string

const (
	EventURLAccessed EventSubject = "url.accessed"
)

func (e EventSubject) IsValid() bool {
	switch e {
	case EventURLAccessed:
		return true
	default:
		return false
	}
}
