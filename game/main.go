package main

import (
	"fmt"
	"syscall/js"

	"github.com/fkmhrk/web-simple-rpg/command"
	"github.com/fkmhrk/web-simple-rpg/model"
)

var state *model.GameState

func initFunc(this js.Value, args []js.Value) interface{} {
	state = model.NewGameState()
	return js.ValueOf(true)
}

func sendCommandFunc(this js.Value, args []js.Value) interface{} {
	out := command.Exec(state, args)
	return js.ValueOf(out)
}

func registerCallbacks() {
	js.Global().Set("init", js.FuncOf(initFunc))
	js.Global().Set("sendCommand", js.FuncOf(sendCommandFunc))
}

func main() {
	fmt.Printf("Start!\n")
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
