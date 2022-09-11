package repository

import (
	"ardamock/src/model"
	"time"

	"gorm.io/gorm"
)

type endpointRepository struct {
	db *gorm.DB
}

type EndpointRepositoryContract interface {
	GetDetail(id int) (*model.Endpoint, error)
	Insert(req model.Endpoint) (model.Endpoint, error)
}

func NewEndpointRepository(db *gorm.DB) EndpointRepositoryContract {
	return &endpointRepository{db: db}
}

func (r *endpointRepository) GetDetail(id int) (*model.Endpoint, error) {
	endpoint := new(model.Endpoint)
	err := r.db.Model(endpoint).
		Where("id", id).
		Find(&endpoint).Error

	return endpoint, err
}

func (r *endpointRepository) Insert(req model.Endpoint) (model.Endpoint, error) {
	var data model.Endpoint

	data.Method = req.Method
	data.Name = req.Name
	data.ExpiredAt = req.ExpiredAt
	data.CreatedAt = time.Now()
	err := r.db.Model(data).Create(&data).Error

	return data, err
}
