package domain

type Notificar interface {
	Notify(message string) error
}
