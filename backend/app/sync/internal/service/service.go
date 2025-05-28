package service

import (
	"context"
	"dex/app/sync/internal/service/corntask"
	"fmt"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Run(ctx context.Context) {
	fmt.Println("solana task run")
	corntask.NewTask().Run()
}
