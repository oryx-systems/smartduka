package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
	"github.com/oryx-systems/smartduka/pkg/smartduka/presentation"
)

const waitSeconds = 30

func main() {
	ctx := context.Background()

	port, err := strconv.Atoi(helpers.MustGetEnvVar(common.PortEnvVarName))
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}
	srv := presentation.PrepareServer(ctx, port, presentation.SmartdukaServiceAllowedOrigins)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			helpers.LogStartupError(ctx, err)
		}
	}()

	// Block until we receive a sigint (CTRL+C) signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*waitSeconds)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until timeout
	err = srv.Shutdown(ctx)
	log.Printf("graceful shutdown started; the timeout is %d secs", waitSeconds)
	if err != nil {
		log.Printf("error during clean shutdown: %s", err)
		os.Exit(-1)
	}
	os.Exit(0)
}
