package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Service struct{}

type Request struct {
	Name string
	Data string
}

type Response struct {
	Success bool
	Message string
	Result  string
}

func (s *Service) GetData(req Request, resp *Response) error {
	log.Printf("GetData called with name: %s", req.Name)
	resp.Success = true
	resp.Message = "Data retrieved successfully"
	resp.Result = "Data for " + req.Name
	return nil
}

func (s *Service) SendData(req Request, resp *Response) error {
	log.Printf("SendData called with name: %s, data: %s", req.Name, req.Data)
	resp.Success = true
	resp.Message = "Data received successfully"
	resp.Result = "Processed " + req.Data
	return nil
}

func main() {
	service := new(Service)
	rpc.Register(service)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("RPC Service listening on :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
