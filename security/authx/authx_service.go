package authx

type Service interface {
	Auth() error
}
