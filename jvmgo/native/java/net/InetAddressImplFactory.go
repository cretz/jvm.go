package io

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_iaif(iaif_isIPv6Supported, "isIPv6Supported", "()Z")
}

func _iaif(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/InetAddressImplFactory", name, desc, method)
}

//static native boolean isIPv6Supported();
// ()Z
func iaif_isIPv6Supported(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(true)
}
