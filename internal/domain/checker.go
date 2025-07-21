package domain

type CheckSaoSilvestre interface {
	Checker() (bool, error)
}
