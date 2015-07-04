package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_cge(cge_initCocoa, "initCocoa", "()V")
	_cge(cge_getMainDisplayID, "getMainDisplayID", "()I")
}

func _cge(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/awt/CGraphicsEnvironment", name, desc, method)
}

func cge_initCocoa(frame *rtda.Frame) {
	//TODO
}

func cge_getMainDisplayID(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
