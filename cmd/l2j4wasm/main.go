package main

import (
	"fmt"
	"syscall/js"

	l2j "github.com/sarkarshuvojit/lomboktojson/pkg"
)

func lombokToJson(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		reportWasmError("Missing argument: lombok string")
		return js.ValueOf("")
	}
	input := args[0].String()
	output, err := safeLombokToJson(input)
	if err != nil {
		reportWasmError(err.Error())
		return js.ValueOf("")
	}
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

func safeLombokToJson(input string) (output string, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("panic while parsing: %v", recovered)
			output = ""
		}
	}()
	jsonStr, err := l2j.LombokToJson(input)
	if err != nil {
		return "", err
	}
	return *jsonStr, nil
}

func reportWasmError(message string) {
	handler := js.Global().Get("onLombokToJsonError")
	if handler.Type() == js.TypeFunction {
		handler.Invoke(js.ValueOf(message))
	}
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
