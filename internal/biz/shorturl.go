package biz

import (
	"context"
	"time"

	v1 "shorturl-server/api/v1"
	"shorturl-server/internal/conf"
	"shorturl-server/internal/provider/kgs"

	kgsapi "github.com/hi-horan/kgs/api/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrHaveNoKeys          = errors.NotFound(v1.ErrorReason_HAVE_NO_KEYS.String(), "have no keys")
	ErrMustHaveOriginalURL = errors.BadRequest(v1.ErrorReason_MUST_HAVE_ORIGINAL_URL.String(), "must have original url")
	ErrInternalError       = errors.InternalServer(v1.ErrorReason_INTERNAL_ERROR.String(), "internal error")
)

type ShortUrl struct {
	OriginalUrl string
	ShortKey    string
	ExpireTime  time.Time
}

type ShortUrlRepo interface {
	CreateShortUrl(context.Context, *ShortUrl) (*ShortUrl, error)
	GetOriginalUrl(context.Context, *ShortUrl) (*ShortUrl, error)
}

type Cache interface {
	Get(context.Context, string) (string, error)

	Set(context.Context, string, string, time.Duration) error
	NewLocker(*LockOpt) Locker
}

type LockOpt struct{}

type Locker interface {
	Lock(context.Context, string, time.Duration) error
	UnLock(context.Context) error
}

type ShortUrlUsecase struct {
	kgsClient *kgs.KGSClient
	suRepo    ShortUrlRepo
	tm        Transaction
	cache     Cache // TODO: 锁 短->长 长->短 共用 如负载过大可考虑分开
	config    *conf.Bootstrap
	log       *log.Helper
}

func NewShortUrlUsecase(repo ShortUrlRepo, tm Transaction, cache Cache, kgsClient *kgs.KGSClient, config *conf.Bootstrap, logger log.Logger) *ShortUrlUsecase {
	return &ShortUrlUsecase{
		kgsClient: kgsClient,
		suRepo:    repo,
		tm:        tm,
		cache:     cache,
		config:    config,
		log:       log.NewHelper(logger),
	}
}

func (uc *ShortUrlUsecase) CreateShortUrl(ctx context.Context, s *ShortUrl) (*ShortUrl, error) {
	l := uc.log.WithContext(ctx)
    l.Infof("CreateShortUrl: %#v", s)

	if len(s.OriginalUrl) == 0 {
		return nil, ErrMustHaveOriginalURL
	}

	shortKey, err := uc.cache.Get(ctx, "ourl:"+s.OriginalUrl)
	if err != nil {
		return nil, ErrInternalError
	}
	if len(shortKey) > 0 {
		l.Debugf("get shorturl from cache. ourl=%s", s.OriginalUrl)
		s.ShortKey = shortKey
		return s, nil
	}

	locker := uc.cache.NewLocker(nil)
	err = locker.Lock(ctx, "l:ourl:"+s.OriginalUrl, time.Second*3)
	if err != nil {
		return nil, ErrInternalError
	}
	defer locker.UnLock(ctx)

	shortKey, err = uc.cache.Get(ctx, "ourl:"+s.OriginalUrl)
	if err != nil {
		return nil, ErrInternalError
	}
	if len(shortKey) > 0 {
		l.Debugf("get shorturl from cache. ourl=%s", s.OriginalUrl)
		s.ShortKey = shortKey
		return s, nil
	}

	reply, err := uc.kgsClient.GetKeys(ctx, &kgsapi.GetKeysRequest{Count: 1}) // TODO: request more once
	if err != nil {
		l.Errorf("failed to call GetKeys. err=%v", err)
		return nil, err
	}
	if len(reply.Keys) == 0 {
		l.Errorf("have no keys")
		return nil, err
	}
	s.ShortKey = reply.Keys[0]

	s, err = uc.suRepo.CreateShortUrl(ctx, s)
	if err != nil {
		return nil, ErrInternalError
	}

	cacheTtl := uc.config.Server.GetCacheShortUrlTime().AsDuration()
	if !s.ExpireTime.IsZero() && s.ExpireTime.Before(time.Now().Local().Add(cacheTtl)) { // url 过期时间过短
		cacheTtl = s.ExpireTime.Sub(time.Now().Local())
	}
	uc.cache.Set(ctx, "skey:"+s.ShortKey, s.OriginalUrl, cacheTtl) // cache set 出错仍继续
	t := uc.config.Server.GetCacheOriginalUrlTime().AsDuration()
	if t < cacheTtl {
		cacheTtl = t
	}
	uc.cache.Set(ctx, "ourl:"+s.OriginalUrl, s.ShortKey, cacheTtl) // cache set 出错仍继续

	l.Infof("CreateShortUrl success. shortUrl=%#v", s)
	return s, nil
}

func (uc *ShortUrlUsecase) GetOriginalUrl(ctx context.Context, s *ShortUrl) (*ShortUrl, error) {
	ourl, _ := uc.cache.Get(ctx, "skey:"+s.ShortKey) // cache出错继续查db
	if ourl == "-1" {
		return nil, ErrHaveNoKeys
	}
	if len(ourl) > 0 {
		s.OriginalUrl = ourl
		return s, nil
	}

	// TODO: 限流

	locker := uc.cache.NewLocker(nil)
	// 防止缓存穿透, 查询db要加锁
	err := locker.Lock(ctx, "l:skey:"+s.ShortKey, 2*time.Second)
	if err != nil {
		return nil, ErrInternalError
	}
	defer locker.UnLock(ctx)

	dbs, err := uc.suRepo.GetOriginalUrl(ctx, s)
	if err == ErrHaveNoKeys || dbs.ExpireTime.Before(time.Now().Local()) {
		// 防止缓存击穿
		uc.cache.Set(ctx, "skey:"+s.ShortKey, "-1", uc.config.Server.CacheInvalidShortUrlTime.AsDuration())
		// cache set 出错仍继续
		return nil, ErrHaveNoKeys
	} else if err != nil {
		return nil, ErrInternalError
	}

	cacheTtl := uc.config.Server.GetCacheShortUrlTime().AsDuration()
	if !dbs.ExpireTime.IsZero() && dbs.ExpireTime.Before(time.Now().Local().Add(cacheTtl)) { // url 过期时间过短
		cacheTtl = dbs.ExpireTime.Sub(time.Now().Local())
	}

	uc.cache.Set(ctx, "skey:"+dbs.ShortKey, dbs.OriginalUrl, cacheTtl) // cache set 出错仍继续
	return dbs, nil
}
