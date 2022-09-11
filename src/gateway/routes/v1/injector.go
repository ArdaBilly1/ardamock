package v1

import (
	v1 "ardamock/src/gateway/handler"
	"ardamock/src/repository"
	"ardamock/src/services"

	"gorm.io/gorm"
)

func EndpointInjector(db *gorm.DB) v1.EndpointHandlerContract {
	repo := repository.NewEndpointRepository(db)
	repoReq := repository.NewRequestRepository(db)
	srv := services.NewEndpointService(repo, repoReq)
	return v1.NewEndpointHandler(srv)
}
