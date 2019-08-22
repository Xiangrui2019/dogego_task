package servers

import "github.com/robfig/cron"

type JobServer struct {
	CronServer *cron.Cron
}

func (server *JobServer) RegisterCronTask() error {
	return nil
}

func (server *JobServer) StartTasks() {

}
