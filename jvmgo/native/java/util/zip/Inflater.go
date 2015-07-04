package zip

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_inflater(inflater_initIDs, "initIDs", "()V")
}

func _inflater(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/util/zip/Inflater", name, desc, method)
}

// private static native void initIDs();
// ()V
func inflater_initIDs(frame *rtda.Frame) {
	// todo
}
