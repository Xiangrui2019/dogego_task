package servers

import (
	"dogego_task/conf"
	"fmt"

	"github.com/robfig/cron"
)

type JobServer struct {
	CronServer *cron.Cron
}

func NewJobServer() *JobServer {
	return &JobServer{
		CronServer: cron.New(),
	}
}

func (server *JobServer) RegisterCronTask() error {
	for x, v := range conf.Conf.Tasks {
		fmt.Println(x, v)
	}

	return nil
}

func (server *JobServer) StartTasks() {

}
