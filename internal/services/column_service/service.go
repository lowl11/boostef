package column_service

type Service struct {
	entity    any
	tableName string
}

func New(entity any) *Service {
	return &Service{
		entity: entity,
	}
}
