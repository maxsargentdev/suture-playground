package coffeedrinker

import (
	"context"
	"go.uber.org/zap"
	"time"
)

type CoffeeDrinker struct {
	Name string
	Log  *zap.SugaredLogger
}

func (service *CoffeeDrinker) Serve(ctx context.Context) error {
	for {

		java := "americano"

		service.Log.Infof("Failed to slurp coffee: %s", java)

		time.Sleep(time.Millisecond)
	}
}
