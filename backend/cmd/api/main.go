package main

import (
	"fmt"

	"github.com/keito-isurugi/next-go-project/internal/infra/db"
	"github.com/keito-isurugi/next-go-project/internal/infra/env"
	"github.com/keito-isurugi/next-go-project/internal/infra/logger"
	"github.com/keito-isurugi/next-go-project/internal/server"
)

func main() {
	ev, err := env.NewValue()
	if err != nil {
		fmt.Println(err.Error())
	}

	zapLogger, err := logger.NewLogger(ev.Debug)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func() { _ = zapLogger.Sync() }()

	dbClient, err := db.NewClient(&ev.DB, zapLogger)
	if err != nil {
		zapLogger.Error(err.Error())
	}

	// awsClient, err := aws.NewAWSSession(ev)
	// if err != nil {
	// 	zapLogger.Error(err.Error())
	// }

	router := server.SetupRouter(ev, dbClient, zapLogger)

	router.Logger.Fatal(router.Start(":" + ev.BeServerPort))

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!!!!!!!")
	// })
	// e.Logger.Fatal(e.Start(":8080"))
}
