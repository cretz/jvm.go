package awt

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _cursor(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}
