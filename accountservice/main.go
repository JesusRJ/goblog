package main

import (
	"fmt"

	"github.com/jesusrj/goblog/accountservice/dbclient"
	"github.com/jesusrj/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Println("Starting", appName)
	initializeBoltClient()
	service.StartWebServer("3000")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}
