package main

/*服务端的方法*/
import (
	"golang.org/x/net/context"
	pb "test-gRPC/friday"
	"fmt"
	"io"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	routeNotes []*pb.UserInfoResponse
}

const (
	port = ":50051"
)

//简单模式
func (this *Server) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	uid := in.GetUid()
	fmt.Println("this is client uid ", uid)
	return &pb.UserInfoResponse{
		Name:  "Jim",
		Age:   18,
		Sex:   0,
		Count: 1000,
	}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//双向流模式
func (this *Server) ChangeUserInfo(stream pb.Data_ChangeUserInfoServer) (error) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			return nil
		}
		if err != nil {
			fmt.Println("ERR", err)
			return err
		}
		fmt.Println("userinfo ", in)
		for _, note := range this.routeNotes {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
