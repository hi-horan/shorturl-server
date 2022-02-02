package data

import (
	"context"

	"shorturl-server/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

)

type dbRepo struct {
	data *Data
	log  *log.Helper
}

func NewShortUrlRepo(data *Data, logger log.Logger) biz.ShortUrlRepo {
	return &dbRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dbRepo) CreateShortUrl(ctx context.Context, s *biz.ShortUrl) (*biz.ShortUrl, error) {
    return r.data.CreateShortUrl(ctx, s)
}

func (r *dbRepo) GetOriginalUrl(ctx context.Context, s *biz.ShortUrl) (*biz.ShortUrl, error) {
    return r.data.GetShortUrl(ctx, s)
}
