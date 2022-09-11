package handler

import (
	"ardamock/src/gateway/handler/request"
	"ardamock/src/services"
	"ardamock/utils/common/constant"
	"ardamock/utils/common/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type endpointHandler struct {
	service services.EndpointServiceContract
}

type EndpointHandlerContract interface {
	GetDetailEndpoint(c echo.Context) error
	InsertData(c echo.Context) error
}

func NewEndpointHandler(srv services.EndpointServiceContract) EndpointHandlerContract {
	return &endpointHandler{service: srv}
}

func (h *endpointHandler) GetDetailEndpoint(c echo.Context) error {
	param := c.Param("id")
	endpointId, _ := strconv.Atoi(param)

	get, err := h.service.GetDetail(endpointId)
	if err != nil || get.ID == 0 {
		return responses.GenerateResponse(
			c,
			http.StatusBadRequest,
			constant.MessageErrorNotFound,
			err,
		)
	}

	return responses.GenerateResponse(
		c,
		http.StatusOK,
		constant.MessageSuccessGetData,
		get,
	)
}

func (h *endpointHandler) InsertData(c echo.Context) error {
	req := new(request.RequestInsertEndpoint)

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return responses.GenerateResponse(
			c,
			http.StatusBadRequest,
			constant.MessageErrorFailedInsert,
			err,
		)
	}

	if err := h.service.Insert(*req); err != nil {
		return responses.GenerateResponse(
			c,
			http.StatusInternalServerError,
			constant.MessageErrorFailedInsert,
			err,
		)
	}

	return responses.GenerateResponse(
		c,
		http.StatusOK,
		constant.MessageSuccessStoreData,
		nil,
	)
}
