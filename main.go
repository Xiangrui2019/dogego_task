package main

import (
	_ "dogego_task/conf"
	"dogego_task/servers"
	"log"
)

func main() {
	server := servers.NewJobServer()

	err := server.RegisterCronTask()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started cron tasks..")
	server.CronServer.Start()

	for {
	}
}
