package service

import (
	"context"
	"time"

	v1 "shorturl-server/api/v1"
	"shorturl-server/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/transport"
)

type ShortUrlService struct {
	v1.UnimplementedShortUrlServer

	uc *biz.ShortUrlUsecase
}

func NewShortUrlService(uc *biz.ShortUrlUsecase) *ShortUrlService {
	return &ShortUrlService{uc: uc}
}

func (s *ShortUrlService) CreateShortUrl(ctx context.Context, in *v1.CreateShortUrlRequest) (*v1.CreateShortUrlReply, error) {
    var et time.Time
    if in.TtlSeconds > 0 {
        et = time.Now().Add(time.Duration(in.TtlSeconds)*time.Second)
    }
	su, err := s.uc.CreateShortUrl(ctx, &biz.ShortUrl{OriginalUrl: in.OriginalUrl, ExpireTime: et})
	if err != nil {
		return nil, err
	}
	return &v1.CreateShortUrlReply{ShortKey: su.ShortKey}, nil
}

func (s *ShortUrlService) GetOriginalUrl(ctx context.Context, in *v1.GetOriginalRequest) (*v1.GetOriginalReply, error) {
    su, err := s.uc.GetOriginalUrl(ctx, &biz.ShortUrl{ShortKey: in.ShortKey})
    if err != nil {
        return nil, err
    }
    return &v1.GetOriginalReply{OriginalUrl: su.OriginalUrl}, nil
}

func (s *ShortUrlService) Redirect(ctx context.Context, in *v1.RedirectRequest) (*v1.RedirectReply, error) {
    su, err := s.uc.GetOriginalUrl(ctx, &biz.ShortUrl{ShortKey: in.ShortKey})
    if err != nil {
        return nil, err
    }
    if tr, ok := transport.FromServerContext(ctx); ok {
		tr.ReplyHeader().Set("Location", su.OriginalUrl)
	}
    err = errors.New(302, "Moved Temporarily", "Moved Temporarily")
    return nil, err
}