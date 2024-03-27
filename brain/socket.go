package brain

import (
	"fmt"
	"log"
	"net"
)

type SocketInterface interface {
	Read() ([]byte, error)
	Write([]byte) (int, error)
	Close() error
}

type Socket struct {
	Socket net.Conn
}

func (s *Socket) Read() ([]byte, error) {
    buffer := make([]byte, 1024)
    n, err := s.Socket.Read(buffer)
    if err != nil {
        log.Printf("Error reading from socket: %v", err)
        return nil, err
    }
    return buffer[:n], nil
}

func (s *Socket) Write(data []byte) (int, error) {
    return s.Socket.Write(data)
}

func (s *Socket) Close() error {
    return s.Socket.Close()
}

func GetConnectionFromListener(port int) SocketInterface {
    listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        log.Fatalf("Failed to start listener: %v", err)
    }
    log.Printf("Server listening on port %v", port)

    conn, err := listener.Accept()
    if err != nil {
        log.Fatalf("Failed to accept connection: %v", err)
    }

    return &Socket{Socket: conn}
}
