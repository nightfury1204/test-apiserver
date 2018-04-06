package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/golang/glog"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/util/logs"
	"github.com/nightfury1204/test-apiserver/commands"
	"github.com/nightfury1204/test-apiserver/pkg/server"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewTryapiServerOptions(os.Stdout, os.Stderr)
	cmd := commands.NewCommandStartServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}