package main

import (
	"os"
	"os/signal"
	"syscall"

	log "logan"

	"./mongo"
	"./rpc"
	"./web"
)

func main() {
	mongo.Start()
	go rpc.Start()

	var stop = make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	go web.Start()
	log.Printf("web server is up")

	sig := <-stop
	log.Printf("web server is down after: %#v", sig)
	web.Stop()
}
