package io

import (
	"net"

	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_sos(sos_init, "init", "()V")
	_sos(sos_socketWrite0, "socketWrite0", "(Ljava/io/FileDescriptor;[BII)V")
}

func _sos(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/SocketOutputStream", name, desc, method)
}

// native java/net/SocketOutputStream~init~()V
func sos_init(frame *rtda.Frame) {

}

//    private native void socketWrite0(FileDescriptor fd, byte[] b, int off,
//                                     int len) throws IOException;
// java/net/SocketOutputStream~socketWrite0~(Ljava/io/FileDescriptor;
func sos_socketWrite0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	fd := vars.GetRef(1)
	b := vars.GetRef(2)
	offset := vars.GetInt(3)
	length := vars.GetInt(4)
	conn := fd.Extra().(net.Conn)
	goBytes := b.GoBytes()
	goBytes = goBytes[offset : offset+length]

	if _, err := conn.Write(goBytes); err != nil {
		frame.Thread().ThrowIOException(err.Error())
	}
}
