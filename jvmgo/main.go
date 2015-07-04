package main

import (
	"os"

	"github.com/cretz/jvm.go/jvmgo/cmdline"
	"github.com/cretz/jvm.go/jvmgo/jvm"
)

func main() {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		jvm.Startup(cmd)
	}
}
