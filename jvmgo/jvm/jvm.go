package jvm

import (
	"os"
	"runtime/pprof"

	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/classpath"
	"github.com/cretz/jvm.go/jvmgo/cmdline"
	"github.com/cretz/jvm.go/jvmgo/jutil"
	"github.com/cretz/jvm.go/jvmgo/jvm/interpreter"
	"github.com/cretz/jvm.go/jvmgo/jvm/keepalive"
	"github.com/cretz/jvm.go/jvmgo/jvm/options"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
	_ "github.com/cretz/jvm.go/jvmgo/native"
)

func Startup(cmd *cmdline.Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	options.InitOptions(cmd.Options())

	cp := classpath.Parse(cmd.Options().Classpath())
	rtc.InitBootLoader(cp)

	mainClassName := jutil.ReplaceAll(cmd.Class(), ".", "/")
	mainThread := createMainThread(mainClassName, cmd.Args())
	interpreter.Loop(mainThread)
	keepalive.KeepAlive()
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := rtc.BootstrapMethod()
	bootArgs := []Any{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
