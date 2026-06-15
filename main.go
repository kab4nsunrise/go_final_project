package main

import (
    "log"
    "todo/pkg/server"
)1

func main() {
    srv := server.New()
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
