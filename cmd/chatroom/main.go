package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/server"
)

var (
	addr = "ChatRoom，start on:2022"
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf(addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
