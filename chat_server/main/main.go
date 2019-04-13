package main

import (
	"time"
	"fmt"
)

func main() {
    fmt.Println("Starting to server...")
	initRedis("localhost:6379", 16, 1024, time.Second*300)
	initUserMgr()
	runServer("0.0.0.0:10000")
}
