package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/trylife/atk/services"
)

var Router = gin.Default()

func init() {
    Router.GET("/app/:appType/:appId/token", services.GetToken)
    Router.GET("/pong", services.Pong)
}
