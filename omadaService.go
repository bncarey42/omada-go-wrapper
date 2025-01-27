package main

type OmadaService[C any, S any, E any] interface {
	Create(entityConfig ...any) error
	GetList(page int32, pageSize int32) ([]S, error)
	GetOne(id string) (E, error)
	Update() error // TODO
	Delete(id string) error
}
