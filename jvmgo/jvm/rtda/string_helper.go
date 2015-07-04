package rtda

import (
	"unicode/utf16"

	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

// todo: is there a better way to create String?
// go string -> java.lang.String
func JString(goStr string) *rtc.Obj {
	internedStr := getInternedString(goStr)
	if internedStr != nil {
		return internedStr
	}

	chars := _stringToUtf16(goStr)
	charArr := rtc.NewCharArray(chars)
	jStr := rtc.BootLoader().JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", charArr)
	return InternString(goStr, jStr)
}

// java.lang.String -> go string
func GoString(jStr *rtc.Obj) string {
	charArr := jStr.GetFieldValue("value", "[C").(*rtc.Obj)
	return _utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func _stringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func _utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
