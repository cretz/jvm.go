package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func _frame(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}
