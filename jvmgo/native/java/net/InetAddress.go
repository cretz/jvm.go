package io

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_ia(ia_init, "init", "()V")
}

func _ia(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/InetAddress", name, desc, method)
}

func ia_init(frame *rtda.Frame) {

}
