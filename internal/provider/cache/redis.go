package cache

import (
	"context"
	"time"

	"shorturl-server/internal/biz"
	"shorturl-server/internal/conf"

    "github.com/bsm/redislock"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
    "github.com/redis/go-redis/extra/redisotel/v9"
)

type cache struct {
    rdb *redis.Client
	log  *log.Helper
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
    l := c.log.WithContext(ctx)
    val, err := c.rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", nil
    } else if err != nil {
        l.Errorf("failed to get cache. key=%s err=%v", key, err)
        return "", err
    }
    return val, nil
}

func (c *cache) Set(ctx context.Context, key, value string, ttl time.Duration) (error) {
    l := c.log.WithContext(ctx)
    err := c.rdb.Set(ctx, key, value, ttl).Err()
    if err != nil {
        l.Errorf("failed to set cache. key=%s val=%s ttl=%v err=%v", key, value, ttl, err)
    }
    return err
}

func (c *cache) NewLocker(opt *biz.LockOpt) biz.Locker {
    return &locker{
        c:c,
        opt: opt,
    }
}

type locker struct {
    c *cache
    l *redislock.Lock
    opt *biz.LockOpt // TODO: no used
}

func (l *locker)Lock(ctx context.Context, key string, ttl time.Duration) (error) {
    lg := l.c.log.WithContext(ctx)
    locker := redislock.New(l.c.rdb)

    var cancel context.CancelFunc
    ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Second*2))
    defer cancel()

    retry := redislock.ExponentialBackoff(time.Millisecond*50, time.Second*2)
    opt := redislock.Options{
        RetryStrategy: retry,
    }

    lock, err := locker.Obtain(ctx, key, ttl, &opt)
    if err != nil {
        lg.Errorf("failed to lock. key=%s err=%v", key, err)
        return err
    }
    l.l = lock
    return nil
}

func (l *locker) UnLock(ctx context.Context) (error) {
    lg := l.c.log.WithContext(ctx)
    err := l.l.Release(ctx)
    if err != nil {
        lg.Errorf("failed to lock. err=%v", err)
    }
    return err
}


func NewRedis(c *conf.Data_Redis, logger log.Logger) (biz.Cache, func (), error) {
    rdb := redis.NewClient(&redis.Options{
		Addr: c.Addr,
		DB:		  0,  // use default DB
	})

    l := log.NewHelper(logger)

    if err := redisotel.InstrumentTracing(rdb); err != nil {
        panic(err)
    }
    if err := redisotel.InstrumentMetrics(rdb); err != nil {
        panic(err)
    }

	cleanup := func() {
		log.NewHelper(logger).Info("closing redis conn")
        rdb.Close()
	}

    return &cache{
        rdb: rdb,
        log: l,
    }, cleanup, nil
}
