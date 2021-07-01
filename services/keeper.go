package services

import (
    "github.com/gin-gonic/gin"
    "github.com/trylife/atk/consts"
    "github.com/trylife/atk/models"
    "github.com/trylife/atk/services/wxmp"
    "github.com/trylife/atk/storage/redis"
    "net/http"
)


func RemoteGetFuc(appType string) models.RemoteGetFunc {
    switch appType {
    case consts.AppWxmp:
        return wxmp.RemoteGet
    }
    return wxmp.RemoteGet
}


func GetToken(c *gin.Context) {
    appType := c.Param("appType")
    appId := c.Param("appId")

    var remoteGetFuc models.RemoteGetFunc

    switch appType {
    case consts.AppWxmp:
        remoteGetFuc = wxmp.RemoteGet
    default:
        c.JSON(http.StatusOK, gin.H{
            "err_code": consts.ErrUndefined,
            "err_msg":  "undefined app type",
        })
        return
    }

    storage := redis.Storage()
    token, err := storage.Get(appType, appId, remoteGetFuc)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "err_code": consts.ErrUndefined,
            "err_msg":  err,
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "err_code": consts.ErrSuc,
        "token":    token,
    })
}

func Pong(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "err_code": 0,
    })
}
