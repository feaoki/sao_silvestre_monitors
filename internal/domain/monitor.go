package domain

type MonitorInscricao interface {
	MonitorarInscricao() (bool, error)
}
