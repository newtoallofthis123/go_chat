package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Addr string
}

func (s *Server) getUserAndMessage(str string) (string, string) {
	split := strings.Split(str, ":")

	return split[0], strings.Join(split[1:], "")
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	log.Default().Println("Connected from:", conn.RemoteAddr().String())

	for {
		buf := make([]byte, 10280)
		n, err := conn.Read(buf)
		if err != nil {
			log.Default().Println("conn.Read err:", err)
			return
		}

		str := string(buf[:n])
		if str == "exit" {
			return
		}

		user, message := s.getUserAndMessage(str)
		log.Default().Printf("user:%s, message:%s\n", user, message)
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Default().Println("net.Listen err:", err)
		return err
	}

	fmt.Println("Listening on address", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Default().Println("listener.Accept err:", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func main() {
	server := Server{Addr: "localhost:8080"}
	server.ListenAndServe()
}
