package settings_test

import (
    "github.com/trylife/atk/settings"
    "os"
    "testing"
)

func TestNew(t *testing.T) {
    pwd, _ := os.Getwd()
    c := settings.NewConf(pwd + "/../app.yaml")
    t.Log(c)
}
