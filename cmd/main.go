package main

import (
	"context"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/f1rdavsi/reporter"
	"github.com/f1rdavsi/reporter/db"
	"github.com/f1rdavsi/reporter/internal/handler"
	"github.com/f1rdavsi/reporter/internal/repository"
	"github.com/f1rdavsi/reporter/internal/service"
	"github.com/f1rdavsi/reporter/logger"
	"github.com/f1rdavsi/reporter/models"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()

	logger.Init()

	db.StartDbConnection()

	_db := db.GetDBConn()
	if err := _db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Category{},
		&models.Transaction{},
	); err != nil {
		logger.Error.Fatal("failed to migrate tables")
	}

	repository := repository.NewRepository(_db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(reporter.Server)
	go func() {
		if err := srv.Run(utils.AppSettings.AppParams.PortRun, handler.InitRoutes()); err != nil {
			logger.Error.Fatal("Error occurred while running http server. Error is: ", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	db.DisconnectDB(_db)
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error.Fatal("Error occurred on server shutting down. Error is: ", err.Error())
		return
	}
}
