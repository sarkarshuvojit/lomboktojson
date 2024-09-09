package main

import (
	"fmt"
	"syscall/js"

	"github.com/sarkarshuvojit/lomboktojson/lib"
)

func lombok2JsonWrapper() js.Func {
	lombok2Jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		in := args[0].String()
		return lib.LombokToJson(in)
	})
	return lombok2Jsonfunc
}

func main() {
	ch := make(chan struct{}, 0)
	fmt.Println("Go Web Assembly")
	js.Global().Set("lombok2json", lombok2JsonWrapper())
	<-ch
}
