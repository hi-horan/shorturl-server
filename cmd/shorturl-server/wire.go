//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"shorturl-server/internal/biz"
	"shorturl-server/internal/conf"
	"shorturl-server/internal/data"
	"shorturl-server/internal/server"
	"shorturl-server/internal/service"
	"shorturl-server/internal/provider"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Bootstrap, *conf.Data, *conf.Data_Redis, *conf.KGS, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, provider.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
