package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5002")
	if err != nil {
		log.Println("failed to resolce tcp address", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("failed to connect to server", err)
	}

	_, err = conn.Write([]byte("hello from " + strconv.Itoa(os.Getpid())))
	if err != nil {
		log.Println(err)
		return
	}

	msg := make([]byte, 1024)
	_, err = conn.Read(msg)
	if err != nil {
		log.Println(err)
		return
	}

	println("message from server:", string(msg))
	conn.Close()
}
