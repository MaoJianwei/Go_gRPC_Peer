package main

import (
	pb "./grpc.maojianwei.com/api"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr = "[::]:9876"
)

var (
	responseCount uint32 = 0
)

type RealBigmaoServer struct {
	pb.UnimplementedBigmaoServer
}

func (s *RealBigmaoServer) QingdaoRequest(ctx context.Context, request *pb.MaoRequestData) (*pb.MaoResponseData, error) {
	responseCount++
	if responseCount % 10000 == 0 {
		log.Printf("Count: %d", responseCount)
	}
	return &pb.MaoResponseData{Count: responseCount, Location: &pb.GPS{Latitude: 40.116 + float64(responseCount),
			Longitude: 116.0987123456789135792468015926 + 2 * float64(responseCount)}}, nil
}

//type RealBigmaoServer struct {
//	pb.UnimplementedBigmaoServer
//}
//
//func (s *RealBigmaoServer) QingdaoRequest(ctx context.Context, request *pb.MaoRequestData) (*pb.MaoResponseData, error) {
//	return &pb.MaoResponseData{Count: responseCount, Location: &pb.GPS{Latitude: 40.116, Longitude: 116.0987123456789}}, nil
//}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("listen fail, %v", err)
	}

	log.Printf("listen ok, %v", listener)

	server := grpc.NewServer()
	log.Printf("server ok, %v", server)

	pb.RegisterBigmaoServer(server, &RealBigmaoServer{})
	if err := server.Serve(listener); err != nil {
		log.Printf("Serve fail, %v", err)
	}
	log.Printf("serve ok, %v", err)
}




























