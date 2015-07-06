package rtda

import (
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	thread       *Thread
	method       *rtc.Method
	localVars    *LocalVars
	operandStack *OperandStack
	maxLocals    uint
	maxStack     uint
	nextPC       int // the next instruction after the call
	onPopAction  func()
}

func newFrame(thread *Thread, method *rtc.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		maxLocals:    method.MaxLocals(),
		maxStack:     method.MaxStack(),
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: NewOperandStack(method.MaxStack()),
	}
}

func (self *Frame) reset(method *rtc.Method) {
	self.method = method
	self.nextPC = 0
	self.lower = nil
	self.onPopAction = nil
	if self.maxLocals > 0 {
		self.localVars.clear()
	}
	if self.maxStack > 0 {
		self.operandStack.Clear()
	}
}

// getters & setters
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *rtc.Method {
	return self.method
}
func (self *Frame) LocalVars() *LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) SetOnPopAction(f func()) {
	self.onPopAction = f
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

// todo
func (self *Frame) ClassLoader() *rtc.ClassLoader {
	return rtc.BootLoader()
}
func (self *Frame) ConstantPool() *rtc.ConstantPool {
	return self.method.ConstantPool()
}
