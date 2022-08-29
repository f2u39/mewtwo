package user

import "net"

type User struct {
	Addr string
	Conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	return &User{
		Addr: addr,
		Conn: conn,
	}
}
