package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// Create new object
type new_ struct {
	Index16Instruction
	class *rtc.Class
}

func (self *new_) Execute(frame *rtda.Frame) {
	if self.class == nil {
		cp := frame.ConstantPool()
		kClass := cp.GetConstant(self.Index).(*rtc.ConstantClass)
		self.class = kClass.Class()
	}

	// init class
	if self.class.InitializationNotStarted() {
		frame.RevertNextPC() // undo new
		frame.Thread().InitClass(self.class)
		return
	}

	ref := self.class.NewObj()
	frame.OperandStack().PushRef(ref)
}
