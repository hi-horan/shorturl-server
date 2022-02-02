package kgs

import (
	"context"
	"time"

	"shorturl-server/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hi-horan/kgs/api/v1"
    grpcx "google.golang.org/grpc"
)

type KGSClient struct {
    v1.KGSClient
	log  *log.Helper
}

func NewKGS(c *conf.KGS, logger log.Logger) (*KGSClient, func (), error) {
    log :=  log.NewHelper(logger)
    // conn, err := grpc.Dial(c.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))

    conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(c.Address),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
		grpc.WithTimeout(2*time.Second),
		// for tracing remote ip recording
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)

	if err != nil {
		log.Fatalf("failed to connect kgs. address=%v err=%w", c.Address, err)
        return nil, nil, err
	}
	client := v1.NewKGSClient(conn)

	cleanup := func() {
		log.Info("closing kgs conn")
        conn.Close()
	}

    return &KGSClient{
        client,
        log,
    }, cleanup, nil
}