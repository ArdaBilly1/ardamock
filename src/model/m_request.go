package model

type BatchRequest []Request

type Request struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	EndpointId int    `json:"endpoint_id"`
	Field      string `json:"field" gorm:"size:100"`
	TypeData   string `json:"type_data" gorm:"size:50"`
	Rule       string `json:"rule" gorm:"size:100"`
}

func (Request) TableName() string {
	return "request"
}
