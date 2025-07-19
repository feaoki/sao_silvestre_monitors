package domain

type Notificar interface {
	Notificar(mensagem string) error
}
