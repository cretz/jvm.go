package io

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_fs(getFileSystem, "getFileSystem", "()Ljava/io/FileSystem;")
}

func _fs(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileSystem", name, desc, method)
}

// public static native FileSystem getFileSystem()
// ()Ljava/io/FileSystem;
func getFileSystem(frame *rtda.Frame) {
	thread := frame.Thread()
	unixFsClass := frame.ClassLoader().LoadClass("java/io/UnixFileSystem")
	if unixFsClass.InitializationNotStarted() {
		frame.SetNextPC(thread.PC()) // undo getFileSystem
		thread.InitClass(unixFsClass)
		return
	}

	stack := frame.OperandStack()
	unixFsObj := unixFsClass.NewObj()
	stack.PushRef(unixFsObj)

	// call <init>
	stack.PushRef(unixFsObj) // this
	constructor := unixFsClass.GetDefaultConstructor()
	thread.InvokeMethod(constructor)
}
