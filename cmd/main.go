package main

import (
	"log"

	"github.com/feaoki/sao-silvestre-watcher/internal/adapters/genai"
	notificador "github.com/feaoki/sao-silvestre-watcher/internal/adapters/print_console"
	"github.com/feaoki/sao-silvestre-watcher/internal/scheduler"
)

func main() {
	log.Println("Iniciando watcher da São Silvestre...")

	checker := genai.NewGenAIChecker()
	notifier := notificador.NewPrintConsole()

	job := &scheduler.WatchJob{
		Checker:   checker,
		Notifier:  notifier,
		Triggered: false,
	}

	scheduler.StartScheduler(job)

	select {} // bloqueia para manter app rodando
}
