package column_service

type Service struct {
	entity any
}

func New(entity any) *Service {
	return &Service{
		entity: entity,
	}
}
