package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test_client/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connectin error:", err)
	}
	defer conn.Close()
	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{"tama"}
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Printf("result: %#v\n", res)
	fmt.Printf(" error: %#v\n", err)

	message2 := &pb.GetMyCatsMessage{TargetCats: []string{"tama", "mike", "john", "mike"}}
	res2, err := client.GetMyCats(context.TODO(), message2)
	fmt.Printf("result: %#v\n", res2)
	for i, res := range res2.MyCat {
		fmt.Printf("result:%d= %#v\n",i, res)
	}
	fmt.Printf(" error: %#v\n", err)

}
