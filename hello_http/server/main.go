package main

import (
	"io"
	"log"
	"net"

	pb "github.com/jergoo/go-grpc-example/proto/hello_http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService 实现服务端接口
var HelloHTTPService = helloHTTPService{}

// SayHello ...
func (h helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

// SayHello ...
func (h helloHTTPService) StreamHello(gs pb.HelloHTTP_StreamHelloServer) error {
	for {
		in, err := gs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		log.Printf("Recv:%s", in.Name)
		gs.Send(&pb.HelloHTTPResponse{Message: "Hello " + in.Name})
	}

	return nil
}
func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloHTTPService
	pb.RegisterHelloHTTPServer(s, HelloHTTPService)

	grpclog.Println("Listen on " + Address)

	s.Serve(listen)
}
