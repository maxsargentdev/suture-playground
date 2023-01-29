package main

import (
	"context"
	"fmt"
	"github.com/thejerf/suture/v4"
	"go.uber.org/zap"
	"suture-testing/coffeedrinker"
	"suture-testing/incrementer"
	"suture-testing/jobrunner"
	"suture-testing/webserver"
)

func main() {
	// Create production logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// Create supervisor
	supervisor := suture.NewSimple("Supervisor")

	// Configure services
	incrementor := &incrementer.IncrementorJob{Log: sugar, Current: 0, Next: make(chan int)}
	drinker := &coffeedrinker.CoffeeDrinker{Log: sugar}
	runner := &jobrunner.JobRunner{Log: sugar}
	web := &webserver.WebServer{Log: sugar}

	// Add services
	supervisor.Add(incrementor)
	supervisor.Add(drinker)
	supervisor.Add(runner)
	supervisor.Add(web)

	// Start service
	ctx, _ := context.WithCancel(context.Background())
	supervisor.ServeBackground(ctx)

	// Misc debugging
	fmt.Println("	Got:", <-incrementor.Next)
	fmt.Println("	Got:", <-incrementor.Next)
	fmt.Println("	Got:", <-incrementor.Next)
	fmt.Println(supervisor.Services())
}
