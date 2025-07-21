package scheduler

import (
	"log"

	"github.com/feaoki/sao-silvestre-watcher/internal/domain"
	"github.com/robfig/cron/v3"
)

type WatchJob struct {
	Checker   domain.MonitorInscricao
	Notifier  domain.Notificar
	Triggered bool
}

func (w *WatchJob) Run() {
	open, err := w.Checker.Check()
	if err != nil {
		log.Printf("Erro ao checar: %v", err)
		return
	}
	if open && !w.Triggered {
		log.Println("Inscrição aberta! Notificando...")
		w.Triggered = true
		_ = w.Notifier.Notify("Inscrições da São Silvestre estão abertas! 🏃")
	}
}

func StartScheduler(job *WatchJob) {
	c := cron.New()
	_, _ = c.AddJob("@every 10m", job)
	c.Start()
}
