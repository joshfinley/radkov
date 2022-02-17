package main

import (
	"context"
	"log"
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/rkpb"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func checkFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func testMain() {
	// dial server
	conn, err := grpc.Dial(
		":1337", grpc.WithTransportCredentials(
			insecure.NewCredentials()))
	checkFatal(err)

	// create stream and its chan
	client := rkpb.NewRadarClient(conn)
	stream, err := client.StreamPlayerPositions(context.Background())
	checkFatal(err)

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

func TestClientStream(t *testing.T) {
	go testMain()
	conn, err := grpc.Dial(":1337", grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	client := rkpb.NewRadarClient(conn)
	stream, err := client.PlayerPositionsStream(context.Background())
	//waitc := make(chan struct{})
	go func() {
		for {
			in, _ := stream.Recv()

			log.Println(in)
		}
	}()

}
