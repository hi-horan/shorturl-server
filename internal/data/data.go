package data

import (
	"context"
	"shorturl-server/internal/biz"
	"shorturl-server/internal/conf"
	"shorturl-server/internal/data/models"
	"time"
    "errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
    "go.opentelemetry.io/otel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "gorm.io/plugin/opentelemetry/tracing"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTransaction, NewShortUrlRepo)

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// Data .
type Data struct {
    db *gorm.DB
	log  *log.Helper
}

type contextTxKey struct{}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	} else {
        return d.db.WithContext(ctx)
    }
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    l:= log.NewHelper(logger)
    db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
    if err != nil {
        l.Errorf("failed to open db. address=%v err=%v", c.Database.Source, err)
        return nil, nil, err
    }
    if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
        dbInstance, _ := db.DB()
        _ = dbInstance.Close()
	}
	return &Data{db:db, log: l}, cleanup, nil
}

func (d *Data) CreateShortUrl(ctx context.Context, in *biz.ShortUrl) (*biz.ShortUrl, error) {
    tr := otel.Tracer("db")
    ctx, span := tr.Start(ctx, "db.CreateShortUrl")
    defer span.End()

    l := d.log.WithContext(ctx)
    now := time.Now().Local()
    var expireTime *time.Time
    if !in.ExpireTime.IsZero() {
        expireTime = &in.ExpireTime
    }
    su := models.ShortUrl{ShortKey: in.ShortKey, OriginalUrl: in.OriginalUrl, CreateTime: now, UpdateTime: now, ExpireTime: expireTime }
    resp := d.DB(ctx).Create(&su)
    if resp.Error != nil {
        l.Errorf("failed to create shorturl(%#v). resp=%#v", su, resp)
        return nil, resp.Error
    } else {
        l.Infof("db create over. shorturl= resp=%#v", su, resp)
        return in, nil
    }
}

func (d *Data) GetShortUrl(ctx context.Context, in *biz.ShortUrl) (*biz.ShortUrl, error) {
    tr := otel.Tracer("db")
    ctx, span := tr.Start(ctx, "db.GetShortUrl")
    defer span.End()

    l := d.log.WithContext(ctx)
    su := models.ShortUrl{}
    resp := d.DB(ctx).Where(&models.ShortUrl{ShortKey: in.ShortKey}).First(&su)
    if resp.Error != nil {
        l.Errorf("failed to create shorturl(%#v). resp=%#v", su, resp)
        if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
            return nil, biz.ErrHaveNoKeys
        }
        return nil, resp.Error
    } else {
        l.Infof("db create over. shorturl= resp=%#v", su, resp)
        if su.ExpireTime != nil {
            in.ExpireTime = *su.ExpireTime
        }
        in.OriginalUrl = su.OriginalUrl
        return in, nil
    }
}