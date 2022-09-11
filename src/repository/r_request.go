package repository

import (
	"ardamock/src/model"

	"gorm.io/gorm"
)

type tableRequestRepositoy struct {
	db *gorm.DB
}

type TableRequestRepositoryContract interface {
	Insert(req []model.Request) error
	GetByEndpointId(endpointId int) (result []model.Request, err error)
}

func NewRequestRepository(db *gorm.DB) TableRequestRepositoryContract {
	return &tableRequestRepositoy{db: db}
}

func (r *tableRequestRepositoy) GetByEndpointId(endpointId int) (result []model.Request, err error) {
	var model model.Request

	err = r.db.Model(model).
		Where("endpoint_id", endpointId).
		Find(&result).
		Error

	return
}

func (r *tableRequestRepositoy) Insert(req []model.Request) error {
	return r.db.Model(req).Create(&req).Error
}
