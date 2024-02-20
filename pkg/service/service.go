package service

import (
	"log"

	"github.com/sylphritz/go-api-server/pkg/db"
	"gorm.io/gorm"
)

type ForeignKeyQuery struct {
	Key   string
	Value any
}

type DatabaseService[T any] interface {
	FindAll(page, perPage int, foreignKey *ForeignKeyQuery) (*[]T, error)
	FindByID(id uint, foreignKey *ForeignKeyQuery) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}

type Service[T any] struct {
	Name string
}

func getDbWhereQuery(foreignKey *ForeignKeyQuery) *gorm.DB {
	if foreignKey.Key != "" {
		return db.DB.Where(foreignKey.Key+" = ?", foreignKey.Value)
	}

	return db.DB
}

const DefaultPage, DefaultPerPage = 1, 10

func NewService[T any](name string) *Service[T] {
	return &Service[T]{Name: name}
}

func (serv *Service[T]) FindAll(page, perPage int, foreignKey *ForeignKeyQuery) (*[]T, error) {
	var results []T

	offset := (page - 1) * perPage

	result := getDbWhereQuery(foreignKey).Offset(offset).Limit(perPage).Find(&results)
	if result.Error != nil {
		log.Printf("An error has occurred while trying to query %v: %v", serv.Name, result.Error)
		return nil, result.Error
	}

	return &results, nil
}

func (serv *Service[T]) FindById(id uint, foreignKey *ForeignKeyQuery) (*T, error) {
	var item T

	result := getDbWhereQuery(foreignKey).First(&item, id)
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

func (serv *Service[T]) Count(foreignKey *ForeignKeyQuery) (int64, error) {
	var model T
	var count int64

	result := getDbWhereQuery(foreignKey).Model(&model).Count(&count)
	if result.Error != nil {
		log.Printf("An error has occurred while trying to count %v: %v", serv.Name, result.Error)
		return count, result.Error
	}

	return count, nil
}
