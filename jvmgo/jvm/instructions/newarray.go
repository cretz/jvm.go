package instructions

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// Create new array
type Newarray struct {
	atype uint8
}

func (self *Newarray) fetchOperands(decoder *InstructionDecoder) {
	self.atype = decoder.readUint8()
}
func (self *Newarray) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		frame.Thread().ThrowNegativeArraySizeException()
		return
	}

	arr := rtc.NewPrimitiveArray(self.atype, uint(count))
	stack.PushRef(arr)
}

func (self *Newarray) ByteSize() int {
	return 2
}

// Create new array of reference
type anewarray struct{ Index16Instruction }

func (self *anewarray) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	kClass := cp.GetConstant(self.Index).(*rtc.ConstantClass)
	componentClass := kClass.Class()

	if componentClass.InitializationNotStarted() {
		thread := frame.Thread()
		frame.SetNextPC(thread.PC()) // undo anewarray
		thread.InitClass(componentClass)
		return
	}

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		frame.Thread().ThrowNegativeArraySizeException()
	} else {
		arr := rtc.NewRefArray(componentClass, uint(count))
		stack.PushRef(arr)
	}

}

// Create new multidimensional array
type Multianewarray struct {
	index      uint16
	dimensions uint8
}

func (self *Multianewarray) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	self.dimensions = decoder.readUint8()
}
func (self *Multianewarray) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	kClass := cp.GetConstant(uint(self.index)).(*rtc.ConstantClass)
	arrClass := kClass.Class()

	stack := frame.OperandStack()
	counts := stack.PopTops(uint(self.dimensions))
	if !_checkCounts(counts) {
		frame.Thread().ThrowNegativeArraySizeException()
	} else {
		arr := _newMultiArray(counts, arrClass)
		stack.PushRef(arr)
	}
}

func (self *Multianewarray) ByteSize() int {
	return 4
}

func _checkCounts(counts []Any) bool {
	for _, c := range counts {
		if c.(int32) < 0 {
			return false
		}
	}
	return true
}

func _newMultiArray(counts []Any, arrClass *rtc.Class) *rtc.Obj {
	count := uint(counts[0].(int32))
	arr := rtc.NewArray(arrClass, count)

	if len(counts) > 1 {
		objs := arr.Refs()
		for i := range objs {
			objs[i] = _newMultiArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
