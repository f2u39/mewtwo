package server

import (
	"fmt"
	"log"
	"mewtwo/goroutines/chat/user"
	"net"
	"sync"
)

type Server struct {
	ip   string
	port int

	// A channel for holds incoming messages
	messages chan []byte

	// A map holds all users in the server
	users map[*user.User]bool

	// A channel for user checkin the server
	checkin chan *user.User

	// A channel for user checkout the server
	checkout chan *user.User
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:    ip,
		port:  port,
		messages: make(chan []byte),
		users: make(map[*user.User]bool),
		checkin: make(chan *user.User),
		checkout: make(chan *user.User),
	}
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	for {
		select {
		case u := <-s.checkin:
			s.users[]
		}

		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.Checkin(conn)
		go s.ReceiveMsg()
		go s.PublishMsg()
	}
}
