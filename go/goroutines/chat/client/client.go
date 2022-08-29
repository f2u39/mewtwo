package main

import (
	"log"
	"net"
	"os"
	"strconv"
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

	sendMsg(conn, "hello from client")

	for {
		if receiveMsg(conn) != nil {
			return
		}
	}
}

func sendMsg(conn *net.TCPConn, msg string) {
	_, err := conn.Write([]byte("hello from " + strconv.Itoa(os.Getpid())))
	if err != nil {
		log.Println(err)
		return
	}
}

func receiveMsg(conn *net.TCPConn) error {
	msg := make([]byte, 1024)
	_, err := conn.Read(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	println(string(msg))
	return nil
}
