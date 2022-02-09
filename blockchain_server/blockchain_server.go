package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

func NewBlockChainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World")
}

func (bsc *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bsc.port)), nil))
}
