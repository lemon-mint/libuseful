package libuseful

import (
	"reflect"
	"unsafe"
)

//go:linkname ProcPin runtime.procPin
//go:linkname ProcUnpin runtime.procUnpin
//go:linkname GetM runtime.getm

func ProcPin() int
func ProcUnpin()
func GetM() unsafe.Pointer

var gType reflect.Type = MustGetType("*runtime.g").Elem()
var offset_g_goid uintptr = MustGetOffset(gType, "goid")
var mType reflect.Type = MustGetType("*runtime.m").Elem()
var offset_m_curg uintptr = MustGetOffset(mType, "curg")
var offset_m_id uintptr = MustGetOffset(mType, "id")
var offset_m_p uintptr = MustGetOffset(mType, "p")
var pType reflect.Type = MustGetType("*runtime.p").Elem()
var offset_p_id uintptr = MustGetOffset(pType, "id")

func GetG() unsafe.Pointer {
	m := GetM()
	return *(*unsafe.Pointer)(unsafe.Add(m, offset_m_curg))
}

func GetP() unsafe.Pointer {
	m := GetM()
	return *(*unsafe.Pointer)(unsafe.Add(m, offset_m_p))
}

func GetMID() int64 {
	m := GetM()
	return *(*int64)(unsafe.Add(m, offset_m_id))
}

func GetGoID() int64 {
	g := GetG()
	return *(*int64)(unsafe.Add(g, offset_g_goid))
}

func GetPID() int {
	p := GetP()
	return int(*(*int32)(unsafe.Add(p, offset_p_id)))
}
