package service

type MessageService interface {
	CountMessages() (int, error)
	CreateMessage(value string) error
}

type Message struct {
	Value string `json:"value"`
}
