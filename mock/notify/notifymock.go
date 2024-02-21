package mock

type Notifier interface {
	Send(message string) error
}

type MessageService struct {
	notifier Notifier
}

func NewMessageService(notifier Notifier) *MessageService {
	return &MessageService{
		notifier: notifier,
	}
}

func (m *MessageService) SendAlert(message string) error {
	return m.notifier.Send(message)
}
