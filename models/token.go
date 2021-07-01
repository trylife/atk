package models

type Token struct {
    AppType     string `json:"app_type"`
    AppId       string `json:"app_id"`
    AccessToken string `json:"access_token"`
    ExpiresIn   int64  `json:"expires_in"`
}

// RemoteGetFunc todo scale param
// RemoteGetFunc uses to get remote access token
type RemoteGetFunc func(appType, appId string) (token Token, err error)
