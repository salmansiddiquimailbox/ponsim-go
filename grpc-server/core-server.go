package grpc_server

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sync"
	pb "openolt"
)

const (
	port = "localhost:50060"
)

func StartGRPCServer(){
	var w sync.WaitGroup
	log.Println("GRPC Server")
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("Erro: %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterOpenoltServer(s, &Server{})

	reflection.Register(s)

	w.Add(1)

	go func(){
		defer w.Done()
		log.Println("Starting gRPC Server")
		er:=s.Server(lis)
		if er!=nil{
		log.Fatalf("Error! Unable to server...")
		}
	}
	w.Wait()
}
