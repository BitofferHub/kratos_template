package task

import (
	"context"
	"fmt"
	pb "github.com/bitstormhub/bitstorm/userX/api/userX/v1"
	"github.com/bitstormhub/bitstorm/userX/internal/conf"
	"github.com/bitstormhub/bitstorm/userX/internal/service"
	"github.com/google/uuid"
	"github.com/google/wire"
	"time"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTaskServer)

type TaskServer struct {
	// 需要什么service, 就修改成自己的service
	service   *service.UserXService
	scheduler *TaskScheduler
}

func (t *TaskServer) Stop(ctx context.Context) error {
	t.scheduler.Stop()
	return nil
}

// 添加Job方法
func (t *TaskServer) NewJobs() []Job {
	return []Job{t.job1, t.job2}
}

// 注入对应service
func NewTaskServer(s *service.UserXService, c *conf.Server) *TaskServer {
	t := &TaskServer{
		service: s,
	}
	conf := c.GetTask()
	t.scheduler = NewScheduler(NewTasks(conf, t.NewJobs()))

	return t
}
func NewTasks(c *conf.Server_TASK, jobs []Job) []*Task {
	var tasks []*Task
	for i, job := range jobs {
		tasks = append(tasks, &Task{
			Name:     c.Tasks[i].Name,
			Type:     c.Tasks[i].Type,
			Schedule: c.Tasks[i].Schedule,
			Handler:  job,
		})
	}

	return tasks
}

// 添加Job方法
func (t *TaskServer) job1() {
	t.service.Cronjob(context.Background(), &pb.CreateUserRequest{
		UserName: uuid.New().String(),
		Pwd:      "12345",
		Sex:      0,
		Age:      20,
		Email:    "12345678@qq.com",
		Contact:  "11",
		Mobile:   "11",
		IdCard:   "11",
	})
	// // 添加一个task
	// t.scheduler.AddTask(Task{
	// 	Name:     "job2",
	// 	Type:     "once",
	// 	NextTime: time.Now().Add(time.Second * 10),
	// 	Handler:  t.job2,
	// })
}

func (t *TaskServer) job2() {
	fmt.Println("定时任务", time.Now())
	next := time.Now().Add(5 * time.Second)
	t.scheduler.AddTask(Task{
		Name:     "job2",
		Type:     "once",
		NextTime: next,
		Handler:  t.job2,
	})
}

func (t *TaskServer) job3() {
	t.service.CreateUser(context.Background(), &pb.CreateUserRequest{
		UserName: uuid.New().String(),
		Pwd:      "2222222",
		Sex:      0,
		Age:      21,
		Email:    "22222@qq.com",
		Contact:  "222222222",
		Mobile:   "222222",
		IdCard:   "222222",
	})
}
