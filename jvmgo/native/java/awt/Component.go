package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _comp(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Component", name, desc, method)
}
