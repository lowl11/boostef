package converter

import (
	"github.com/lowl11/boostef/data/interfaces/imigrate"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/ef_core"
	"github.com/lowl11/boostef/internal/migrator"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/ctypes"
	"github.com/lowl11/flex"
	"reflect"
	"strings"
)

func (converter *Converter) Entity() imigrate.Entity {
	/*
		what we are looking for?
		1. table name +
		2. columns ~
			2.1. column name +
			2.2. column type +
			2.3. is it nullable? +
			2.4. is it has default value? +
			2.5. is it primary key? +
			2.6. is it foreign key? +
	*/

	var tableName string
	partitionColumns := make([]string, 0)

	flexStruct := flex.Struct(converter.entity)
	fields := flexStruct.Fields()
	for _, field := range fields {
		efTags := flex.Field(field).Tag("ef")
		if len(efTags) > 0 {
			for _, tag := range efTags {
				if strings.Contains(tag, "table") {
					_, tableName, _ = strings.Cut(tag, ":")
				} else if strings.Contains(tag, "part") {
					_, partitionBy, found := strings.Cut(tag, ":")
					if found {
						partitionColumns = strings.Split(partitionBy, ",")
					}
				}
			}
		}
	}

	if tableName == "" {
		panic("table name not found. Name: " + converter.name)
	}

	rowFields := flexStruct.FieldsRow()

	columns := make([]iquery.Column, 0, len(rowFields))
	for _, field := range rowFields {
		flexField := flex.Field(field)
		tags := flexField.Tag("db")
		if len(tags) == 0 {
			panic("no 'db' tag for: " + field.Name)
		}

		var defaultValue string
		var isPrimaryKey bool
		var foreign string
		var isUnique bool

		dt := convertDataType(flexField.Type(), flexField.Tag("ef"))
		if dt == nil {
			continue
		}

		efTags := flexField.Tag("ef")
		for _, tag := range efTags {
			if tag == "pk" {
				isPrimaryKey = true
			} else if tag == "unique" {
				isUnique = true
			} else if strings.Contains(tag, "default:") {
				_, defaultValue, _ = strings.Cut(tag, ":")
			} else if strings.Contains(tag, "fk:") {
				_, foreign, _ = strings.Cut(tag, ":")
			}
		}

		dt.Default(defaultValue)
		if isPrimaryKey {
			dt.Primary()
		}

		if len(foreign) > 0 {
			dt.Foreign(foreign)
		}

		if isUnique && !isPrimaryKey {
			dt.Unique()
		}

		columns = append(columns, builder.
			Column(tags[0]).
			DataType(dt))
	}

	return migrator.
		NewEntity(ef_core.Get().Schema(), tableName, converter.name).
		Columns(columns...).
		PartitionColumns(partitionColumns...)
}

func convertDataType(t reflect.Type, efTags []string) iquery.DataType {
	for _, tag := range efTags {
		if strings.Contains(tag, "custom") {
			_, customType, found := strings.Cut(tag, ":")
			if !found {
				continue
			}

			if len(customType) == 0 {
				continue
			}

			custom := ctypes.Custom(customType)
			isPtr := flex.Type(t).IsPtr()
			if !isPtr {
				custom.NotNull()
			}

			return custom
		}
	}

	flexType := flex.Type(t)
	isPtr := flexType.IsPtr()
	t = flexType.Unwrap()

	var dt iquery.DataType

	switch t.String() {
	case "time.Time":
		dt = ctypes.Timestamp()
	case "uuid.UUID":
		dt = ctypes.Uuid()
	}

	if dt == nil && flexType.IsBytes() {
		dt = ctypes.JsonB()
	}

	if dt != nil {
		if !isPtr {
			dt.NotNull()
		}
		return dt
	}

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		dt = ctypes.Integer()
	case reflect.String:
		if len(efTags) > 0 && efTags[0] == "jsonb" {
			dt = ctypes.JsonB()
		} else {
			dt = ctypes.Text()
		}
	case reflect.Bool:
		dt = ctypes.Boolean()
	case reflect.Float32, reflect.Float64:
		dt = ctypes.Real()
	default:
		return nil
	}

	if !isPtr {
		return dt.NotNull()
	}

	return dt
}

func convertCustomType(custom string) (iquery.DataType, bool) {
	if len(custom) == 0 {
		return nil, false
	}

	ctypes.Custom(custom)

	return nil, false
}
