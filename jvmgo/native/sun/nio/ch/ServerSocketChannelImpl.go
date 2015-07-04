package ch

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _ssci(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
