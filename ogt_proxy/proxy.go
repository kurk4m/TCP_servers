package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	s.ln = ln

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}

		fmt.Println("New conection: ", conn.RemoteAddr())

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(src net.Conn) {
	dst, err := net.Dial("tcp", "https://login.kural.pl/")
	if err != nil {
		log.Fatalln("Unable to connect to target server")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}

	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	server := NewServer(":3000")
	log.Fatal(server.Start())
}
