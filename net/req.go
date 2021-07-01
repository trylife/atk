package net

import (
    "io/ioutil"
    "net/http"
)

// HttpGet simple http get request
func HttpGet(url string) (body []byte, e error) {
    resp, err := http.Get(url)
    if err != nil {
        return body, err
    }
    defer resp.Body.Close()
    body, err = ioutil.ReadAll(resp.Body)

    if err != nil {
        return body, err
    }

    return body, nil
}