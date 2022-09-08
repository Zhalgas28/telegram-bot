package service

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Entity {
	return allEntities
}

func (s *Service) Get(idx int) (*Entity, error) {
	if idx >= len(allEntities) || idx < 0 {
		return nil, fmt.Errorf("entity with id %d doesn't exists", idx)
	}
	return &allEntities[idx], nil
}

func (s *Service) Create(entity Entity) {
	allEntities = append(allEntities, entity)
}

func (s *Service) Delete(idx int) error {
	if len(allEntities) <= idx {
		return fmt.Errorf("entity with id %d doesn't exists", idx)
	}
	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
	return nil
}

func (s *Service) Update(idx int, entity Entity) error {
	if len(allEntities) <= idx {
		return fmt.Errorf("entity with id %d doesn't exists", idx)
	}
	allEntities[idx] = entity
	return nil
}
