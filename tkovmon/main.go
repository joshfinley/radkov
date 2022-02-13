package main

import (
	"context"
	"log"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/rkpb"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CheckFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// dial server
	conn, err := grpc.Dial(
		":1337", grpc.WithTransportCredentials(
			insecure.NewCredentials()))
	CheckFatal(err)

	// create stream and its chan
	client := rkpb.NewServerClient(conn)
	stream, err := client.StreamPlayerPositions(context.Background())
	CheckFatal(err)

	pch := make(chan [][]byte)
	go tarkov.MonitorGame(pch, &tarkov.TarkovOffsets)

	// send
	for d := range pch {
		post := rkpb.PlayerPositions{
			RawVectors: d,
		}

		if err := stream.Send(&post); err != nil {
			log.Fatalf("failed sending: %v", err)
		}

		log.Printf("%d vectors sent", len(post.RawVectors))
	}
}
