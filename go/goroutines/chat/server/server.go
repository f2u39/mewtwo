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

	s.Broadcast("welcome" + u.Addr)
	fmt.Println("UserMap", s.Users)
}

func (s *Server) ReceiveMsg() {
	msg := <-s.MsgCh
	s.Broadcast(msg)
}

func (s *Server) Broadcast(msg string) {
	b := []byte(msg)

	s.Lock.Lock()
	for _, user := range s.Users {
		user.Conn.Write(b)
	}
	s.Lock.Unlock()
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen failed...")
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.checkin(conn)

		go s.ReceiveMsg()
	}
}
