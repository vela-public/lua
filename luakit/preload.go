package luakit

import "github.com/vela-public/lua"

type Preloader interface {
	Set(string, lua.LValue)
	SetGlobal(string, lua.LValue)
}

func PreloadModule(L *lua.LState) {
	kv := lua.NewUserKV()
	kv.Set("map", lua.NewFunction(NewMapL))
	kv.Set("slice", lua.NewFunction(NewSliceL))
	kv.Set("fmt", lua.NewFunction(NewFmtL))
	kv.Set("pretty", lua.NewFunction(PrettyJsonL))
	L.SetGlobal("luakit", kv)
}

func Preload(ip Preloader) {
	kv := lua.NewUserKV()
	kv.Set("map", lua.NewFunction(NewMapL))
	kv.Set("slice", lua.NewFunction(NewSliceL))
	kv.Set("fmt", lua.NewFunction(NewFmtL))
	kv.Set("pretty", lua.NewFunction(PrettyJsonL))
	ip.Set("luakit", kv)
}
