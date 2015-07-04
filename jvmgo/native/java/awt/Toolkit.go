package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _tk(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Toolkit", name, desc, method)
}
