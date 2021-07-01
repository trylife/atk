package wxmp

import (
    "encoding/json"
    "errors"
    "github.com/trylife/atk/consts"
    "github.com/trylife/atk/models"
    "github.com/trylife/atk/net"
    "github.com/trylife/atk/settings"
)

const (
    MpBaseUrl = "https://api.weixin.qq.com/cgi-bin"
)

const (
    mpErrSuc = 0
)

type mpAccessToken struct {
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
    ErrCode     int    `json:"errcode"`
    ErrMsg      string `json:"errmsg"`
}

func RemoteGet(appType, appId string) (token models.Token, err error) {
    c := settings.Conf()
    mp, ok := c.GetApp(appType + appId)
    if !ok {
        return token, errors.New("appId not exists")
    }

    url := MpBaseUrl + "/token?grant_type=client_credential&appid=" + mp.AppId + "&secret=" + mp.Secret
    body, err := net.HttpGet(url)
    if err != nil {
        return token, err
    }

    var ac mpAccessToken
    err = json.Unmarshal(body, &ac)
    if err != nil {
        return token, err
    }

    if ac.ErrCode != mpErrSuc {
        return token, errors.New("remote code " + string(body))
    }

    token.AppType = consts.AppWxmp
    token.AccessToken = ac.AccessToken
    token.AppId = appId
    token.ExpiresIn = ac.ExpiresIn

    return token, nil
}
