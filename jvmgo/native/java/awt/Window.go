package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func _window(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Window", name, desc, method)
}
