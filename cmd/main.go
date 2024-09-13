package main

import (
	"github.com/vela-public/lua"
	"github.com/vela-public/lua/luakit"
)

func main() {

	file := "cmd/vela.lua"

	co := lua.NewState()
	defer co.Close()
	luakit.PreloadModule(co)

	if err := co.DoFile(file); err != nil {
		panic(err)
	}

}
