package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke interface method
type Invokeinterface struct {
	index uint
	// count uint8
	// zero uint8

	// optimization
	kMethodRef   *rtc.ConstantInterfaceMethodref
	argSlotCount uint
}

func (self *Invokeinterface) fetchOperands(decoder *InstructionDecoder) {
	self.index = uint(decoder.readUint16())
	decoder.readUint8() // count
	decoder.readUint8() // must be 0
}

func (self *Invokeinterface) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.index).(*rtc.ConstantInterfaceMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		panic("NPE") // todo
	}

	method := self.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread().InvokeMethod(method)
}

func (self *Invokeinterface) ByteSize() int {
	return 5
}