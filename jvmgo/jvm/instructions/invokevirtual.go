package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke instance method; dispatch based on class
type Invokevirtual struct {
	Index16Instruction
	kMethodRef   *rtc.ConstantMethodref
	argSlotCount uint
}

func (self *Invokevirtual) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.Index).(*rtc.ConstantMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	method := self.kMethodRef.GetVirtualMethod(ref)
	frame.Thread().InvokeMethod(method)
}
