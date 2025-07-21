package scheduler

import (
	"fmt"

	"github.com/feaoki/sao-silvestre-watcher/internal/domain"
	"github.com/robfig/cron/v3"
)

type WatchJob struct {
	Checker   domain.CheckSaoSilvestre
	Notifier  domain.Notificar
	Triggered bool
}

func (w *WatchJob) Run() {
	open, err := w.Checker.Checker()
	if err != nil {
		fmt.Printf("Erro ao checar: %v", err)
		return
	}
	if open && !w.Triggered {
		fmt.Println("Inscrição aberta! Notificando...")
		w.Triggered = true
		_ = w.Notifier.Notify("Inscrições da São Silvestre estão abertas! 🏃")
	}
}

func StartScheduler(job *WatchJob) {
	c := cron.New()
	_, _ = c.AddJob("@every 1m", job)
	c.Start()
}
