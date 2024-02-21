package mock

// 這個系統將允許不同的消息發送方式（如電子郵件、短信或推送通知）來實現同一接口
// MessageService 依賴於一個 Notifier 接口，這個接口的實現決定了消息的具體發送方式。
// 創建一個MessageService 實例時，你會將一個具體的 Notifier（例如 EmailNotifier、SmsNotifier 或 PushNotifier）注入到 MessageService。
// 這樣，MessageService 就可以使用這個注入的 Notifier 來發送消息，而不用關心消息是如何發送的。

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

