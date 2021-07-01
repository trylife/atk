package schedule

import (
    "fmt"
    "github.com/robfig/cron/v3"
    "github.com/trylife/atk/services"
    "github.com/trylife/atk/settings"
    "github.com/trylife/atk/storage/redis"
)

func Start() {
    // cron
    c := cron.New()
    entryId, err := c.AddFunc("@every 1m", func() {
        for _, app := range settings.Conf().AppMap {
            token, err := redis.Storage().Get(app.AppType, app.AppId, services.RemoteGetFuc(app.AppType))
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println("1min:" + token.AccessToken)
        }
    })
    if err != nil {
        panic(err)
    }
    c.Start()
    fmt.Println("entryId: ", entryId)
}
