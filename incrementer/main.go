package incrementer

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const (
	JobLimit = 2
)

type IncrementorJob struct {
	Log     *zap.SugaredLogger
	Current int
	Next    chan int
}

func (service *IncrementorJob) Serve(ctx context.Context) error {
	for {
		select {
		case service.Next <- service.Current + 1:
			service.Current++
			time.Sleep(time.Second)
			if service.Current >= JobLimit {
				fmt.Println("Stopping the service")
			}
		}
	}
}
