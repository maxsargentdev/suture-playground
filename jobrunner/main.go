package jobrunner

import (
	"context"
	"go.uber.org/zap"
	"time"
)

type JobRunner struct {
	Name string
	Log  *zap.SugaredLogger
}

func (service *JobRunner) Serve(ctx context.Context) error {
	for {

		job := "00001-batch"

		service.Log.Infof("Failed to run job: %s", job)

		time.Sleep(time.Millisecond)
	}
}
