package solana

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type SolanaTask struct {
}

func NewSolanaTask() *SolanaTask {
	return &SolanaTask{}
}

func RegisterTask(c *cron.Cron) {
	task := NewSolanaTask()
	c.AddJob("@every 1s", task)
}

func (etherum *SolanaTask) Run() {
	fmt.Println("solana task run")
}
