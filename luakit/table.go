package luakit

import "github.com/vela-public/lua"

func IntOr(tab *lua.LTable, key string, d int) int {
	lv := tab.RawGetString(key)

	switch lv.Type() {
	case lua.LTNumber:
		return int(lv.(lua.LNumber))
	case lua.LTInt:
		return int(lv.(lua.LInt))
	default:
		return d
	}
}

func StringOr(tab *lua.LTable, key string, d string) string {
	lv := tab.RawGetString(key)
	if lv == lua.LNil {
		return d
	}

	return lv.String()
}
