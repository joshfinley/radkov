package main

import (
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/joshfinley/radkov/pkg/rkpb"
	"github.com/joshfinley/radkov/radarsrv/radarapp"
	"google.golang.org/grpc"
)

func TestMain(t *testing.T) {
	// initialize global state store
	radarapp.GlobalGameState.Init()

	// create listener
	lis, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	rkpb.RegisterRadarServer(s, radarapp.RadarServer{})

	// create http server
	// TODO: Development only code - should only use RPC services in final thing
	router := radarapp.NewRadarService()
	log.Println("starting web service on port 80")
	go http.ListenAndServe(":80", &router)
	//

	// start the server
	log.Println("starting rpc service on port 1337")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
