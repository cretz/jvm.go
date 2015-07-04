package atomic

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_al(VMSupportsCS8, "VMSupportsCS8", "()Z")
}

func _al(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/util/concurrent/atomic/AtomicLong", name, desc, method)
}

// private static native boolean VMSupportsCS8();
// ()Z
func VMSupportsCS8(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushBoolean(false) // todo sync/atomic
}
