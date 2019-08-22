package servers

import (
	"dogego_task/conf"
	"fmt"
	"time"

	_ "github.com/levigross/grequests"

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

func (server *JobServer) Run(name string, callback_url string) {
	from := time.Now().UnixNano()
	to := time.Now().UnixNano()

	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}

func (server *JobServer) RegisterCronTask() error {
	for _, v := range conf.Conf.Tasks {
		server.CronServer.AddFunc(v.Time, func() {})
	}

	return nil
}
