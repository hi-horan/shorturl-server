module shorturl-server

go 1.16

require (
	github.com/bsm/redislock v0.9.1
	github.com/go-kratos/kratos/v2 v2.5.4
	github.com/google/wire v0.5.0
	github.com/hi-horan/kgs v0.0.1
	github.com/prometheus/client_golang v1.14.0
	github.com/redis/go-redis/extra/redisotel/v9 v9.0.2
	github.com/redis/go-redis/v9 v9.0.2
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.uber.org/automaxprocs v1.5.1
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
	gorm.io/plugin/opentelemetry v0.1.1
)

replace github.com/hi-horan/kgs => ../kgs
