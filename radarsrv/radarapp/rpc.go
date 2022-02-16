package radarapp

import (
	"io"
	"log"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/rkpb"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

type RadarServer struct {
	rkpb.UnsafeRadarServer
}

func (s RadarServer) StreamPlayerPositions(
	srv rkpb.Radar_StreamPlayerPositionsServer) error {
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

		// copy the memory once and then forget about it
		go GlobalGameState.SetPlayerPositions(vecs)

		res := rkpb.Response{
			Ok: true,
		}
		if err := srv.Send(&res); err != nil {
			log.Printf("send error: %v", err)
		}
	}
}
