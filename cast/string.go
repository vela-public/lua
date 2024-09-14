package cast

import (
	"github.com/vela-public/lua"
	"reflect"
	"unsafe"
)

func S2B(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}

func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func S2L(s string) lua.LString {
	return lua.LString(s)
}

func B2L(b []byte) lua.LString {
	return *(*lua.LString)(unsafe.Pointer(&b))
}
