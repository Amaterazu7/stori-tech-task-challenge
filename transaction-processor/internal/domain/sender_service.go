package domain

type SenderService interface {
	SendMessage() (int, error)
}
