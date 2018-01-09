package main

import (
	"log"
	"strconv"

	pb "github.com/jergoo/go-grpc-example/proto/hello_http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main2() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHTTPClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Message)
}
func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHTTPClient(conn)

	stream, err := c.StreamHello(context.Background())
	if err != nil {
		log.Printf("failed to call: %v", err)
		return
	}

	var i int64
	for {
		stream.Send(&pb.HelloHTTPRequest{Name: strconv.FormatInt(i, 10)})
		if err != nil {
			log.Printf("failed to send: %v", err)
			break
		}
		reply, err := stream.Recv()
		if err != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("Greeting: %s", reply.Message)
		i++
	}
}

// OR: curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
