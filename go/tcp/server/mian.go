package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const port = 5002

func main() {
	lis, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf("listen port %d failed, %s", port, err)
		os.Exit(1)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	var (
		buf = make([]byte, 512)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

RECV_LOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		if data == "" {
			return
		}

		switch err {
		case io.EOF:
			break RECV_LOOP

		case nil:
			log.Println("Client → Server:", data)
			break RECV_LOOP

		default:
			log.Fatal("receive message from client failed", err)
			return
		}
	}

	var m string = "Hello from server!"

	b := []byte(m)
	w.Write(b)
	w.Flush()
	log.Println("Server → Client:", m)
}
