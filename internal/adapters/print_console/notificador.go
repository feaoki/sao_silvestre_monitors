package notificador

import (
	"fmt"

	"github.com/feaoki/sao-silvestre-watcher/internal/domain"
)

type PrintConsole struct{}

func NewPrintConsole() domain.Notificar {
	return &PrintConsole{}
}

func (p *PrintConsole) Notify(message string) error {
	fmt.Println("Mensagem para o console:", message)
	return nil
}
