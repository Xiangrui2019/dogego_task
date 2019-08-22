package main

import (
	_ "dogego_task/conf"
	"dogego_task/servers"
)

func main() {
	server := servers.NewJobServer()
	server.RegisterCronTask()
}
