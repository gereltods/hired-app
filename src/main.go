package main

import (
	"fasthttp/app"
	"flag"
	"fmt"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	fmt.Println("Hello go")
	app.StartApplication()
}
