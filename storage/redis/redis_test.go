package redis_test

import (
    "github.com/trylife/atk/models"
    "github.com/trylife/atk/storage/redis"
    "testing"
)

func TestClient(t *testing.T) {
    storage := redis.Storage()
    err := storage.Set(models.Token{
        AppType:     "wechat_mini_program",
        AppId:       "wx123456789",
        AccessToken: "234567890-",
        ExpiresIn:   7200,
    })

    t.Log("set ", err)

    if err != nil {
        t.Error(err)
    }

    m := func(appType, appId string) (models.Token, error) {
        token := models.Token{
            AppType:     appType,
            AppId:       appId,
            AccessToken: appId,
            ExpiresIn:   7200,
        }
        return token, nil
    }

    appType := "test_type"
    appId := "test_xxxx"
    token, err := storage.Get(appType, appId, m)

    tl, _ := m(appType, appId)
    t.Log("m", tl)
    t.Log("set by ", token)

    if err != nil {
        t.Error(err)
    }

    token, err = storage.Get("wechat_mini_program", "wx123456789", nil)
    t.Log(token)
    if err != nil {
        t.Error(err)
    }

}
