package provider

import (
	"shorturl-server/internal/provider/kgs"
	"shorturl-server/internal/provider/cache"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(kgs.NewKGS, cache.NewRedis)
