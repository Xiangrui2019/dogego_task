package main

import (
	_ "dogego_task/conf"
	"dogego_task/servers"
	"log"
)

func main() {
	server := servers.NewJobServer()
	server.RegisterCronTask()

	log.Println("Started cron tasks..")
	server.CronServer.Start()
	for {
	}
}
