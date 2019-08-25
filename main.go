package main

import (
	_ "dogego_task/conf"
	"dogego_task/servers"
	"log"
)

func main() {
	server := servers.NewJobServer()
	channel := make(chan int)

	err := server.RegisterCronTask()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started cron tasks..")
	server.CronServer.Start()

	<-channel
}
