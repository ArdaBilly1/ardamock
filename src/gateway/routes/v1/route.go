package v1

import (
	"ardamock/utils/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo says: im fine :)")
	})

	v1 := e.Group("/v1")

	endpointHandler := EndpointInjector(config.MysqlDB)
	endpoint := v1.Group("/endpoint")
	endpoint.GET("/:id", endpointHandler.GetDetailEndpoint)
	endpoint.POST("", endpointHandler.InsertData)
}
