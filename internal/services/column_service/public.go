package column_service

import (
	"errors"
	"github.com/lowl11/boostef/data/models"
	"github.com/lowl11/flex"
	"reflect"
)

func (service *Service) CheckType() error {
	if obj := flex.Object(service.entity); !obj.IsStruct() {
		return errors.New("Given entity is not struct but: " + obj.Type().Kind().String())
	}

	return nil
}

func (service *Service) GetColumns(customType ...reflect.Type) []models.Column {
	var s flex.RStruct
	if len(customType) > 0 {
		s = flex.Struct(reflect.New(customType[0]).Interface())
	} else {
		s = flex.Struct(service.entity)
	}

	rowFields := s.FieldsRow()
	fields := make([]flex.RField, 0, len(rowFields))
	for _, field := range rowFields {
		fields = append(fields, flex.Field(field))
	}

	columns := make([]models.Column, 0, len(fields))
	for _, field := range fields {
		if !field.IsPublic() {
			continue
		}

		if tableNames := field.Tag("table"); len(tableNames) > 0 {
			service.tableName = tableNames[0]
			continue
		}

		fieldType := flex.Type(field.Type())
		if fieldType.IsStruct() && !fieldType.IsTime() {
			columns = append(columns, service.GetColumns(field.Type())...)
			continue
		}

		dbTagValue := field.Tag("db")
		if len(dbTagValue) == 0 || dbTagValue[0] == "" {
			continue
		}

		columns = append(columns, models.Column{
			Name:     dbTagValue[0],
			DataType: defineType(field.Type()),
			EfTags:   field.Tag("ef"),
		})
	}

	return columns
}

func (service *Service) TableName() string {
	return service.tableName
}
