package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"myProtobuf/simplerpc"
	"time"
)

var list = []string{"shindou", "中文", "kafka", "otter", "golang"}

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	log.Printf("[grpc client] dial to server success.sleep 10 senconds and restart server")
	time.Sleep(10 * time.Second)
	grpcclient := simplerpc.NewGreeterClient(conn)
	for _, v := range list {
		go runClient(grpcclient, v)
	}
	select {}
}

func runClient(grpcclient simplerpc.GreeterClient, name string) {
	for {
		reply, err := grpcclient.SayHello(context.TODO(), &simplerpc.HelloRequest{Name: name})
		if err != nil {
			log.Printf("[grpc client] grpc call error: %#v\n", err)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			continue
		}
		log.Printf("[grpc client] received grpc server reply: %#v\n", reply.Message)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
