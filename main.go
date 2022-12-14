package main

import (
	v1 "ardamock/src/gateway/routes/v1"
	"ardamock/utils/config"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var serviceConfig config.Service

	service := serviceConfig.SetConfig()
	timezone, _ := time.Now().Zone()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Validator = &config.Validator{Validator: validator.New()}

	v1.Route(e)

	e.Server.Addr = service.Addr
	e.Logger.Printf("Starting server at %s | Timezone: %s", service.Addr, timezone)
	e.Logger.Fatal(e.Server.ListenAndServe())

	// start service
	go func() {
		if err := e.Start(":" + os.Getenv("APP_ADDRESS")); err != nil {
			e.Logger.Info("Shutting down a server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
