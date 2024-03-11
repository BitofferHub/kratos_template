//go:build wireinject

// The build tag makes sure the stub is not built in the final build.
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
	"github.com/google/wire"
)

// wireApp
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: wireApp init kratos application.
//	@param *conf.Server
//	@param *conf.Data
//	@return *kratos.App
//	@return func()
//	@return error
func wireApp(*conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		interfaces.ProviderSet,
		task.ProviderSet,
		newApp))
}
