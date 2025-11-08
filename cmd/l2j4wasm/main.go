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

func beautifyLombok(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return js.ValueOf("")
	}

	indent := 2
	if len(args) > 1 {
		if provided := args[1].Int(); provided > 0 {
			indent = provided
		}
	}

	formatted, err := l2j.Beautify(args[0].String(), indent)
	if err != nil {
		return js.ValueOf(args[0].String())
	}

	return js.ValueOf(formatted)
}

func registerCallbacks() {
	js.Global().Set("lombokToJson", js.FuncOf(lombokToJson))
	js.Global().Set("beautifyLombok", js.FuncOf(beautifyLombok))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
