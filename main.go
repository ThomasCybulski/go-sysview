package main

import (
	"flag"

	astilectron "go-sysview/pkg/gui"
	webserver "go-sysview/pkg/http"
)

var (
	port  string
	debug bool
)

// @title go-sysview
// @version 1.0
// @description API collection

// @BasePath /api/v1
func main() {
	flag.StringVar(&port, "port", "7099", "api server port")
	flag.BoolVar(&debug, "debug", false, "enables the debug mode")
	flag.Parse()

	go webserver.Start(port, debug)
	astilectron.New(port, debug)
}
