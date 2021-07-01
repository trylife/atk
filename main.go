package main

import (
    "github.com/trylife/atk/routes"
    "github.com/trylife/atk/schedule"
)

func main() {

    schedule.Start()

    err := routes.Router.Run()
    if err != nil {
        panic(err)
    }
}
