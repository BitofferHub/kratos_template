package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"

	pb "github.com/bitstormhub/bitstorm/userX/api/userX/v1"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	//	callGRPC()
	callHTTP()
	// callGRPCDiscover()
}

// callGRPC
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:  just a demo for rpc call without discover
func callGRPC() {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:6001"),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserXClient(conn)
	reply, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{UserName: "niuniu"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc] CreateUser reply %+v\n", reply)

}

// callGRPCDiscover
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: just a demo for rpc call with discovery
func callGRPCDiscover() {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	endpoint := "discovery:///user-svr"
	conn, err := transgrpc.DialInsecure(context.Background(), transgrpc.WithEndpoint(endpoint), transgrpc.WithDiscovery(dis))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	cli := pb.NewUserXClient(conn)
	reply, err := cli.CreateUser(context.Background(), &pb.CreateUserRequest{UserName: "niuniu"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc] CreateUser reply %+v\n", reply)

}

// callHTTP
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: just a demo for http call with discovery
func callHTTP() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)
	endpoint := "discovery:///user-svr"
	connHTTP, err := http.NewClient(
		context.Background(),
		http.WithEndpoint(endpoint),
		http.WithDiscovery(dis),
		http.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connHTTP.Close()

	httpClient := pb.NewUserXHTTPClient(connHTTP)
	fmt.Printf("before call\n")
	reply, err := httpClient.GetUser(context.Background(), &pb.GetUserRequest{UserId: 1})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("[http] GetUser %+v\n", reply)
	time.Sleep(10 * time.Second)

}
