package user

import "net"

type User struct {
	Addr    string
	Conn    net.Conn
	MsgChan chan []byte
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	return &User{
		Conn:    conn,
		Addr:    addr,
		MsgChan: make(chan []byte),
	}
}
