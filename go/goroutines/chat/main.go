package main

import "mewtwo/goroutines/chat/server"

const (
	ip   = "127.0.0.1"
	port = 5002
)

func main() {
	srv := server.NewServer(ip, port)
	srv.Start()
}
