package main

/*客户端方法*/

import (
	"golang.org/x/net/context"
	pb "test-gRPC/friday"
	"fmt"
	"io"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

//简单模式
func GetUserInfo(client pb.DataClient, info *pb.UserInfoRequest) {
	req, err := client.GetUserInfo(context.Background(), info)
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	fmt.Println("userinfo is ", req.GetAge(), req.GetCount(), req.GetName(), req.GetSex())
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDataClient(conn)
	req, err := c.GetUserInfo(context.Background(), &pb.UserInfoRequest{Uid: 123})
	if err != nil {
		fmt.Println("Could not create Customer: %v", err)
	}
	fmt.Println("userinfo is ", req.GetAge(), req.GetCount(), req.GetName(), req.GetSex())
}

//双向流模式
func ChangeUserInfo(client pb.DataClient) {
	notes := []*pb.UserInfoResponse{
		{Name: "jim", Age: 18, Sex: 2, Count: 100},
		{Name: "Tom", Age: 20, Sex: 1, Count: 666},
	}
	stream, err := client.ChangeUserInfo(context.Background())
	if err != nil {
		fmt.Println("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				fmt.Println("read done ")
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println("Failed to receive a note : %v", err)
			}
			fmt.Println("Got message %s at point(%d, %d)", in.Count, in.Sex, in.Age, in.Name)
		}
	}()
	fmt.Println("notes", notes)
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			fmt.Println("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
