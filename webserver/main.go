package webserver

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

const (
	JobLimit = 2
)

type WebServer struct {
	Log *zap.SugaredLogger
}

func (service *WebServer) Serve(ctx context.Context) error {
	defer ctx.Done()

	mux := http.NewServeMux()

	service.Log.Infof("Starting webserver")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		ctx.Done()
		return err
	}

	return nil
}
