package settings

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "os"
)

var c = new(Config)
var appsMap = make(map[string]App)

func init() {
    c = NewConf()
}

func Conf() *Config {
    return c
}

type Config struct {
    Redis  Redis `yaml:"redis"`
    Apps   []App `yaml:"apps"`
    AppMap map[string]App
}

type Redis struct {
    Addr     string `yaml:"addr"`
    Password string `yaml:"password"`
    Db       int    `yaml:"db"`
}

// App i.e Wechat Mini Program
type App struct {
    AppType string `yaml:"app_type"`
    AppId   string `yaml:"app_id"`
    Secret  string `yaml:"secret"`
    Comment string `yaml:"comment"`
}

func NewConf(specConfFile ...string) *Config {
    err := c.getConf(specConfFile...)
    if err != nil {
        panic(err)
    }

    c.AppMap = appsMap
    for _, app := range c.Apps {
        c.AppMap[app.AppType+app.AppId] = app
    }

    return c
}

func (c *Config) GetApp(appId string) (w App, ok bool) {
    if appId == "" {
        return w, false
    }

    w, ok = c.AppMap[appId]
    return w, ok
}

func (c *Config) getConf(specConfFile ...string) error {
    var err error
    var confFile string
    userHomeDir, _ := os.UserHomeDir()
    pwd, _ := os.Getwd()
    confFiles := []string{
        userHomeDir + "/app.yaml",
        "./app.yaml",
        pwd + "/../app.yaml",
        pwd + "/../../app.yaml",
    }

    if len(specConfFile) > 0 {
        confFiles = append(specConfFile, confFiles...)
    }

    for _, confFile = range confFiles {
        if _, err = os.Stat(confFile); err == nil {
            break
        }
    }

    if err != nil {
        panic(err)
    }

    data, err := ioutil.ReadFile(confFile)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(data, &c)
    if err != nil {
        panic(err)
    }

    return nil
}
