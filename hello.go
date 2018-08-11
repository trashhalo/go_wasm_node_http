package main

import (
	"fmt"
	"syscall/js"
)

var headers js.Value

var (
	jsFun            = js.Global().Get("Function")
	natveRequestFunc = nativeRequest()
)

func init() {
	headers := js.Global().Get("Object").New()
	headers.Set("Content-Type", "text/plain")
}

func nativeRequest() js.Value {
	return jsFun.New("res", "headers", `
		res.writeHead(200, headers)
		res.write("Hello World")
		res.end("\n")	
	`)
}

func main() {
	require := js.Global().Get("require")
	http := require.Invoke("http")
	app := http.Get("Server").New()
	app.Call("on", "request", js.NewCallback(request))
	app.Call("listen", 3000, js.NewCallback(listen))
	wait()
}

func request(args []js.Value) {
	res := args[1]
	natveRequestFunc.Invoke(res, headers)
}

func listen(args []js.Value) {
	fmt.Println("listening on port 3000")
}

func wait() {
	done := make(chan bool)
	js.Global().Get("process").Call("on", "SIGTERM", js.NewCallback(func(args []js.Value) {
		done <- true
	}))
	<-done
}
