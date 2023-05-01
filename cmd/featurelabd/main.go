package main

import (
	"flag"
	"fmt"
	"github.com/torresjeff/featurelabd"
	"github.com/torresjeff/featurelabd/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {

	var featureLabUrl string
	var port uint
	var ttlMinutes uint
	var cleanUpIntervalMinutes uint

	flag.StringVar(&featureLabUrl, "featurelab-url", "http://localhost:3000", "the url of the Feature Lab Server. Default: http://localhost:3000")
	flag.UintVar(&port, "port", 43743, "the port to listen on. Default: 43743")
	flag.UintVar(&ttlMinutes, "ttlMinutes", 10, "the ttl for the cache in minutes. Default: 10")
	flag.UintVar(&cleanUpIntervalMinutes, "clean-up-interval", 30, "the interval to clean up the cache in seconds. Default: 30")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("featurelab daemon failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFeatureLabServer(grpcServer,
		featurelabd.NewFeatureLabService(featurelabd.NewCacheableFeatureLabClient(featureLabUrl,
			time.Duration(ttlMinutes)*time.Minute,
			time.Duration(cleanUpIntervalMinutes)*time.Minute),
			"FeatureLab", "Test"))

	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}
