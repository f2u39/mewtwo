package server

import (
	"fmt"
	"log"
	"net"
	"os/user"
)

type Server struct {
	Ip   string
	Port int

	Users map[string]*user.User
	MsgCh chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

func (s *Server) checkin(conn net.Conn) {
	fmt.Println("welcome", conn.RemoteAddr().String())
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
	}
}
