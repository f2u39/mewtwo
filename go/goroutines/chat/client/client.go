package main

import (
	"log"
	"net"
)

const (
	ip   = "127.0.0.1"
	port = ":5002"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ip+port)
	if err != nil {
		log.Println("failed to resolce tcp address", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("failed to connect to server", err)
		return
	}

	conn.SetKeepAlive(false)
}
