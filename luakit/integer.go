package luakit

import "github.com/vela-public/lua"

func NewIntL(L *lua.LState) int {
	n := L.IsInt(1)
	L.Push(lua.LInt(n))
	return 1
}
