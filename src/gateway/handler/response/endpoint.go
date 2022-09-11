package response

type DetailReponse struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	RequestField []RequestField `json:"request_field"`
}

type RequestField struct {
	Name     string `json:"name" validate:"required"`
	TypeData string `json:"type_data"`
	Rule     string `json:"rule"`
}
