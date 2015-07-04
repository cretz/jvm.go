package native

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	//rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

type NativeMethod func(frame *rtda.Frame)
