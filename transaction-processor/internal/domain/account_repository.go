package domain

type AccountRepository interface {
	FindById() error
}
