package request

type RequestInsertEndpoint struct {
	Method       string              `json:"method" validate:"required"`
	Name         string              `json:"name" validate:"required"`
	ExpiredHours int                 `json:"expired_hours"`
	Request      []RequestBodyInsert `json:"request"`
}

type RequestBodyInsert struct {
	Name     string `json:"name"`
	TypeData string `json:"type_data"`
	Rule     string `json:"rule"`
}
