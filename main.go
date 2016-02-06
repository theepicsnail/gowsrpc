package main

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"golang.org/x/net/websocket"
)

// This example demonstrates a trivial echo server.
func main() {
	adder := &Adder{0}

	// Reset the counter every 30 seconds
	go func() {
		c := time.Tick(30 * time.Second)
		for _ = range c {
			adder.Reset()
		}
	}()

	// register our adder (adds the exposed methods)
	// set the http server to use /rpc as the websocket endpoint
	rpc.Register(adder)
	http.Handle("/rpc", websocket.Handler(func(ws *websocket.Conn) {
		jsonrpc.ServeConn(ws)
	}))

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
