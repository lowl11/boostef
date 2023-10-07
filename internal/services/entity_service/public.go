package entity_service

import "github.com/lowl11/boostef/data/models"

func (service *Service) Columns() []models.Column {
	return service.columns
}

func (service *Service) TableName() string {
	return service.tableName
}
