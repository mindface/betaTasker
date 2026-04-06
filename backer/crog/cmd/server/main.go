package main

import (
    "log"
    "repo/crog/internal/router"
)

func main() {
    r := router.New()
    log.Fatal(r.Run(":8082"))
}