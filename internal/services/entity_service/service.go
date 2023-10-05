package entity_service

import (
	"github.com/lowl11/boostef/data/models"
	"github.com/lowl11/boostef/internal/services/column_service"
)

type Service struct {
	columns []models.Column
}

func New(entity any) *Service {
	fx := column_service.New(entity)
	if err := fx.CheckType(); err != nil {
		panic(err)
	}

	return &Service{
		columns: fx.GetColumns(),
	}
}
