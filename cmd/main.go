package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"vertexUP/initialize"
	"vertexUP/models"
	"vertexUP/pkg/handler"
	"vertexUP/pkg/repository"
	"vertexUP/pkg/service"
	"vertexUP/pkg/usecase"
	"vertexUP/server"
)

func init() {
	logrus.Info("start init server")
	if err := initialize.LoadConfiguration(); err != nil {
		logrus.Fatal(err.Error())
	}
	if err := initialize.RunLogger(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func main() {
	logrus.Info("start server")

	var serverInstance server.Server

	businessDatabase := repository.NewBusinessDatabase(initialize.Env, &initialize.Config.Database)

	repositorySources := repository.Sources{
		BusinessDB: businessDatabase,
	}
	repos := repository.NewRepository(&repositorySources)
	services := service.NewService(repos, &initialize.Config, initialize.Env)
	usecases := usecase.NewUsecase(services)
	handlers := handler.NewHandler(usecases, services)

	go runServer(serverInstance, handlers, initialize.Env, initialize.Config.Server)
	runChannelStopServer()

	serverInstance.Shutdown(businessDatabase, context.Background())
}

func runServer(server server.Server, handlers *handler.Handler, env *models.Environment, cfg models.ServerConfig) {
	ginEngine := handlers.InitHTTPRoutes(env)
	if err := server.Run(cfg.Port, ginEngine); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", nil, err.Error())
	}
}
func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
