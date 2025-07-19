package main

import (
	"log"

	"github.com/feaoki/sao-silvestre-watcher/internal/adapters/notifier"
	"github.com/feaoki/sao-silvestre-watcher/internal/adapters/scraper"
	"github.com/feaoki/sao-silvestre-watcher/internal/scheduler"
)

func main() {
	log.Println("Iniciando watcher da São Silvestre...")

	checker := scraper.NewCollyChecker("https://www.saosilvestre.com.br/")
	notifier := notifier.NewTelegramNotifier("SEU_BOT_TOKEN", "SEU_CHAT_ID")

	job := &scheduler.WatchJob{
		Checker:   checker,
		Notifier:  notifier,
		Triggered: false,
	}

	scheduler.StartScheduler(job)

	select {} // bloqueia para manter app rodando
}
