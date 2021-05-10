package main

import (
	"context"
	"github.com/sergio-vaz-abreu/star-wars/application"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	err := run()
	if err != nil {
		logrus.WithError(err).Fatal("failed running application")
	}
}

func run() error {
	applicationConfig, err := application.LoadConfigFromEnv()
	if err != nil {
		return err
	}
	logrus.Info("starting application")
	return start(applicationConfig)
}

func start(applicationConfig application.Config) error {
	app, err := application.Load(applicationConfig)
	if err != nil {
		return err
	}
	appErr := app.Run()
	ctx := gracefullyShutdown()
	defer app.Shutdown()
	select {
	case err := <-appErr:
		return err
	case <-ctx.Done():
		return nil
	}
}

func gracefullyShutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		logrus.Info("gracefully shutdown")
		cancel()
	}()
	return ctx
}
