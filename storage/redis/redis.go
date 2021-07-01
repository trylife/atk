package redis

import (
    "context"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "github.com/trylife/atk/models"
    "github.com/trylife/atk/settings"
    "time"
)

var ctx = context.Background()
var o = new(op)

type op struct {
    client *redis.Client
}

func init() {
    o = newClient(settings.Conf())
}

func Storage() *op {
    return o
}

func newClient(settings *settings.Config) *op {
    rdb := redis.NewClient(&redis.Options{
        Addr:     settings.Redis.Addr,
        Password: settings.Redis.Password,
        DB:       settings.Redis.Db,
    })
    o.client = rdb
    rdb.PoolStats()
    return o
}

func (o *op) Set(token models.Token) error {
    key := o.genKey(token)
    val, err := json.Marshal(token)
    if err != nil {
        return err
    }
    err = o.client.Set(ctx, key, val, time.Duration(token.ExpiresIn)*time.Second).Err()
    return err
}

func (o *op) Get(appType, appId string, remoteGet models.RemoteGetFunc) (models.Token, error) {
    token := models.Token{
        AppType: appType,
        AppId:   appId,
    }
    key := o.genKey(token)
    val, err := o.client.Get(ctx, key).Bytes()

    if err == redis.Nil {
        token, err = o.remoteGetAndSet(appType, appId, remoteGet)
        return token, err
    } else if err != nil {
        return token, err
    }

    err = json.Unmarshal(val, &token)
    if err != nil {
        return token, err
    }
    ex := o.client.TTL(ctx, key).Val()

    token.ExpiresIn = ex.Milliseconds() / 1000

    // 5min force refresh
    if token.ExpiresIn < 300 {
        token, err = o.remoteGetAndSet(appType, appId, remoteGet)
        return token, err
    }

    return token, nil
}

func (o *op) remoteGetAndSet(appType, appId string, remoteGet models.RemoteGetFunc) (token models.Token, err error) {
    token, err = remoteGet(appType, appId)

    if err != nil {
        return token, err
    }

    err = o.Set(token)
    if err != nil {
        return token, err
    }
    return token, nil
}

func (o *op) genKey(token models.Token) string {
    return token.AppType + ":" + token.AppId
}
