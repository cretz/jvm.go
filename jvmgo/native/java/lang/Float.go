package lang

import (
	"math"

	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_float(floatToRawIntBits, "floatToRawIntBits", "(F)I")
	_float(intBitsToFloat, "intBitsToFloat", "(I)F")
}

func _float(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Float", name, desc, method)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtda.Frame) {
	vars := frame.LocalVars()
	value := vars.GetFloat(0)
	bits := math.Float32bits(value)

	stack := frame.OperandStack()
	stack.PushInt(int32(bits)) // todo
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtda.Frame) {
	vars := frame.LocalVars()
	bits := vars.GetInt(0)
	value := math.Float32frombits(uint32(bits))

	stack := frame.OperandStack()
	stack.PushFloat(value)
}
