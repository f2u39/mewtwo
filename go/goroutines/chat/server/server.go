package server

import (
	"fmt"
	"log"
	"mewtwo/goroutines/chat/user"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	Lock  sync.RWMutex
	Users map[string]*user.User
	MsgCh chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:    ip,
		Port:  port,
		Users: make(map[string]*user.User),
		MsgCh: make(chan string),
	}
}

func (s *Server) checkin(conn net.Conn) {
	addr := conn.RemoteAddr().String()

	u := user.NewUser(conn)

	s.Lock.Lock()
	s.Users[addr] = u
	s.Lock.Unlock()

	s.Broadcast("welcome", u)
	fmt.Println("UserMap", s.Users)
}

func (s *Server) ReceiveMsg() {
	for {
		msg := <-s.MsgCh

		s.Lock.Lock()
		for _, cli := range s.Users {
			cli.MsgCh <- msg
		}
		s.Lock.Unlock()
	}
}

func (s *Server) Broadcast(msg string, user *user.User) {
	m := user.Addr + ": " + msg
	s.MsgCh <- m
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen failed...")
	}
	defer lis.Close()

	go s.ReceiveMsg()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.checkin(conn)
	}
}
