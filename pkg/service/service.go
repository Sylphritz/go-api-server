package service

import (
	"log"

	"github.com/sylphritz/go-api-server/pkg/db"
)

type DatabaseService[T any] interface {
	FindAll(page, perPage int) (*[]T, error)
	FindByID(id uint) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}

type Service[T any] struct {
	Name string
}

const DefaultPerPage = 10

func NewService[T any](name string) *Service[T] {
	return &Service[T]{Name: name}
}

func (serv *Service[T]) FindAll(page, perPage int) (*[]T, error) {
	var results []T

	offset := (page - 1) * perPage

	result := db.DB.Offset(offset).Limit(perPage).Find(&results)
	if result.Error != nil {
		log.Printf("An error has occurred while trying to query %v: %v", serv.Name, result.Error)
		return nil, result.Error
	}

	return &results, nil
}

func (serv *Service[T]) FindById(id uint) (*T, error) {
	var item T

	result := db.DB.First(&item, id)
	if result.Error != nil {
		log.Printf("An error has occurred while trying to query %v: %v", serv.Name, result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (serv *Service[T]) Create(entity *T) error {
	return db.DB.Create(entity).Error
}

func (serv *Service[T]) Update(entity *T) error {
	return db.DB.Save(entity).Error
}

func (serv *Service[T]) Delete(id uint) error {
	// The model is needed so the DB knows which table to perform delete
	var entity T

	return db.DB.Delete(&entity, id).Error
}
