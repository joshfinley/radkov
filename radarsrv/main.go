package main

import (
	"io"
	"log"
	"net"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/rkpb"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
	"google.golang.org/grpc"
)

type server struct {
	rkpb.UnsafeServerServer
}

func (s server) StreamPlayerPositions(
	srv rkpb.Server_StreamPlayerPositionsServer) error {
	//
	log.Println("starting streaming receiver")
	ctx := srv.Context()

	for {
		// exit if context is done
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from the stream
		post, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
		}

		if err != nil {
			log.Printf("receive error: %v", err)
			continue
		}

		vecs := post.RawVectors
		log.Println("first vector received:",
			unity.UnmarshalVec2(vecs[0]))
		res := rkpb.Response{
			Ok: true,
		}
		if err := srv.Send(&res); err != nil {
			log.Printf("send error: %v", err)
		}
	}
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	rkpb.RegisterServerServer(s, &server{})

	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
