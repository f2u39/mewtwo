package user

import "net"

type User struct {
	Addr string
	Ch   chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	return &User{
		Addr: addr,
		Ch:   make(chan string),
		conn: conn,
	}
}
