package main

import (
	"docker-go/application/db"
	"docker-go/application/route"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db.InitRedis()
	defer db.Close()
	fmt.Println("Registering routes")
	route.RegisterRoute()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	fmt.Println("Shutting down gracefully...")

}
