package servers

import (
	"dogego_task/conf"
	"fmt"
	"net/http"
	"time"

	"github.com/levigross/grequests"
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
	resp, err := grequests.Get(callback_url, nil)
	to := time.Now().UnixNano()

	if err != nil {
		fmt.Printf("任务 %s 运行失败: %dms\n", name, (to-from)/int64(time.Millisecond))
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("任务 %s 运行失败: %dms\n", name, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("任务 %s 运行成功: %dms\n", name, (to-from)/int64(time.Millisecond))
	}
}

func (server *JobServer) RegisterCronTask() error {
	for _, v := range conf.Conf.Tasks {
		server.CronServer.AddFunc(v.Time, func() { server.Run(v.Name, v.CallbackURL) })
	}

	return nil
}
