package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain Server: ")
}

func main() {
	fmt.Println("test")
	// defaultの　ポート指定
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockChainServer(uint16(*port))
	app.Run()
}
