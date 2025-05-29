package etherum

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type EtherumTask struct {
}

func NewEtherumTask() *EtherumTask {
	return &EtherumTask{}
}

func RegisterTask(c *cron.Cron) {
	task := NewEtherumTask()
	c.AddJob("@every 1s", task)

	eventPairCreateTask := NewEventPairCreate()
	c.AddJob("@every 2s", eventPairCreateTask)
}

func (etherum *EtherumTask) Run() {
	fmt.Println("etherum task run")
}
