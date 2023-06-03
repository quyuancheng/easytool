package main

import (
	"os"
	"os/signal"
	"syscall"

	"tool/cmd/tool/config"
	"tool/pkg/controller"
	"tool/pkg/models"
	"tool/pkg/usecase"
	restfulutil "tool/pkg/utils/restful"

	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Parse():
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	switch config.C.LogLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	service := restfulutil.NewService("api", "v1", config.C.ListenHost, config.C.ListenPort, restfulutil.APIForUI, config.C.EnableAccessLog)
	//instance := db.GetDbInst()

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	shellCtrl := controller.NewShellController()
	uiCtrl := controller.NewIndexCtrl()
	objects := []*inject.Object{
		//{Value: instance.DB()},
		{Value: models.NewShellRepo()},
		{Value: usecase.NewShellUcase()},
		{Value: shellCtrl},
	}

	if err := g.Provide(objects...); err != nil {
		logrus.Fatalf("provide objects to the Graph: %v", err)
	}
	if err := g.Populate(); err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
	service.Add(shellCtrl, uiCtrl)
	errChan := make(chan error, 1)
	service.Run(errChan)

	// step finally: listen Signal
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-errChan:
		logrus.Errorf("Received collapse error %s, exiting gracefully...", err.Error())
	case <-term:
		logrus.Warn("Received SIGTERM, exiting gracefully...")
	}
	logrus.Info("See you next time!")
}
