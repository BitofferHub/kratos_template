// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/bitstormhub/bitstorm/userX/internal/biz"
	"github.com/bitstormhub/bitstorm/userX/internal/conf"
	"github.com/bitstormhub/bitstorm/userX/internal/data"
	"github.com/bitstormhub/bitstorm/userX/internal/interfaces"
	"github.com/bitstormhub/bitstorm/userX/internal/server"
	"github.com/bitstormhub/bitstorm/userX/internal/service"
	"github.com/bitstormhub/bitstorm/userX/internal/task"
	"github.com/go-kratos/kratos/v2"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: wireApp init kratos application.
//	@param *conf.Server
//	@param *conf.Data
//	@return *kratos.App
//	@return func()
//	@return error
func wireApp(confServer *conf.Server, confData *conf.Data) (*kratos.App, func(), error) {
	db := data.NewDatabase(confData)
	client := data.NewCache(confData)
	dataData := data.NewData(db, client)
	userRepo := data.NewUserXRepo(dataData)
	transaction := data.NewTransaction(dataData)
	userXUseCase := biz.NewUserXUseCase(userRepo, transaction)
	userXService := service.NewUserXService(userXUseCase)
	grpcServer := server.NewGRPCServer(confServer, userXService)
	handler := interfaces.NewHandler(userXService)
	httpServer := server.NewHTTPServer(confServer, handler)
	taskServer := task.NewTaskServer(userXService, confServer)
	app := newApp(grpcServer, httpServer, taskServer)
	return app, func() {
	}, nil
}
