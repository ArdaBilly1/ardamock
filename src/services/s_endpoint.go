package services

import (
	"ardamock/src/gateway/handler/request"
	"ardamock/src/gateway/handler/response"
	"ardamock/src/model"
	"ardamock/src/repository"
	"ardamock/utils/common/constant"
	"fmt"
	"time"
)

type endpointService struct {
	repository        repository.EndpointRepositoryContract
	requestRepository repository.TableRequestRepositoryContract
}

type EndpointServiceContract interface {
	GetDetail(id int) (response.DetailReponse, error)
	Insert(req request.RequestInsertEndpoint) error
}

func NewEndpointService(
	repo repository.EndpointRepositoryContract,
	repoReq repository.TableRequestRepositoryContract,
) EndpointServiceContract {
	return &endpointService{
		repository:        repo,
		requestRepository: repoReq,
	}
}

const DefaultTimeExpired = 3600

func (s *endpointService) GetDetail(id int) (response.DetailReponse, error) {
	var resp response.DetailReponse

	data, err := s.repository.GetDetail(id)
	if err != nil {
		return resp, err
	}

	if data.ID == 0 {
		return resp, fmt.Errorf(constant.MessageErrorNotFound)
	}

	resp.ID = data.ID
	resp.Method = data.Method
	resp.Name = data.Name

	fieldReq, err := s.requestRepository.GetByEndpointId(data.ID)
	if err != nil {
		return resp, fmt.Errorf(constant.MessageErrorNotFound)
	}

	for _, r := range fieldReq {
		var temp response.RequestField
		temp.Name = r.Field
		temp.Rule = r.Rule
		temp.TypeData = r.TypeData

		resp.RequestField = append(resp.RequestField, temp)
	}

	return resp, err
}

func (s *endpointService) Insert(req request.RequestInsertEndpoint) error {
	mapping := mappingInsert(req)
	data, err := s.repository.Insert(mapping)
	if err != nil {
		return err
	}

	mappingTableRequest := mappingTableRequestInsert(data.ID, req.Request)
	err = s.requestRepository.Insert(mappingTableRequest)

	return err
}

func mappingInsert(req request.RequestInsertEndpoint) model.Endpoint {
	var data model.Endpoint
	data.Method = req.Method
	data.Name = req.Name
	data.ExpiredAt = time.Now().Add(time.Minute * DefaultTimeExpired)

	return data
}

func mappingTableRequestInsert(idEndpoint int, req []request.RequestBodyInsert) model.BatchRequest {
	var data model.BatchRequest

	for _, r := range req {
		var datas model.Request
		datas.EndpointId = idEndpoint
		datas.Field = r.Name
		datas.Rule = r.Rule
		datas.TypeData = r.TypeData

		data = append(data, datas)
	}

	return data
}
