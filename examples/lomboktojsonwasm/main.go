package main

import (
	"syscall/js"

	l2j "github.com/sarkarshuvojit/lomboktojson/pkg"
)

func lombokToJson(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return js.ValueOf("Missing argument: lombok string")
	}
	input := args[0].String()
	println("Input", input)
	jsonStr, err := l2j.LombokToJson(input)
	if err != nil {
		return js.ValueOf("{}")
	}
	output := *jsonStr
	return js.ValueOf(output)

}

func registerCallbacks() {
	js.Global().Set("lombokToJson", js.FuncOf(lombokToJson))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
