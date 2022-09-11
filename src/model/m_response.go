package model

const (
	TypeResponseSuccess                   = 1
	TypeResponseFailedBadRequest          = 2
	TypeResponseFailedInternalServerError = 3
)

type Response struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	EndpointId int    `json:"endpoint_id"`
	Type       int    `json:"type" gorm:"size:5"`
	Detail     string `json:"detail" gorm:"type:json"`
}

func (*Response) TableName() string {
	return "response"
}
