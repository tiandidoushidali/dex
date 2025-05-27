package service

import (
	"context"
	"fmt"
	"github.com/tiandidoushidali/dex-dodo/backend/app/sync/internal/service/corntask"
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
