package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke a class (static) method
type invokestatic struct {
	Index16Instruction
	method *rtc.Method
}

func (self *invokestatic) Execute(frame *rtda.Frame) {
	if self.method == nil {
		cp := frame.Method().Class().ConstantPool()
		k := cp.GetConstant(self.index)
		if kMethodRef, ok := k.(*rtc.ConstantMethodref); ok {
			self.method = kMethodRef.StaticMethod()
		} else {
			self.method = k.(*rtc.ConstantInterfaceMethodref).StaticMethod()
		}
	}

	// init class
	class := self.method.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	frame.Thread().InvokeMethod(self.method)
}
