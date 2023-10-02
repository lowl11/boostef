package crud

import "github.com/lowl11/boostef/internal/repositories/repository"

type Crud struct {
	repository.Repository
}

func New() *Crud {
	return &Crud{}
}
