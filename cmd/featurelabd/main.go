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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var featureLabUrl string
	var port uint
	var ttlMinutes uint
	var cleanUpIntervalMinutes uint

	flag.StringVar(&featureLabUrl, "featureLabUrl", "http://localhost:3000", "the url of the Feature Lab Server. Default: http://localhost:3000")
	flag.UintVar(&port, "port", 43743, "the port for the daemon to listen on. Default: 43743")
	flag.UintVar(&ttlMinutes, "ttlMinutes", 10, "the ttl for the cache in minutes. The Feature Lab daemon will automatically fetch features at intervals specified by this argument. Default: 10")
	flag.UintVar(&cleanUpIntervalMinutes, "cleanUpInterval", 30, "the interval to clean up the cache in minutes. Default: 30")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("featurelab daemon failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	featureLabClient := featurelabd.NewCacheableFeatureLabClient(featureLabUrl,
		time.Duration(ttlMinutes)*time.Minute,
		time.Duration(cleanUpIntervalMinutes)*time.Minute)
	pb.RegisterFeatureLabServer(grpcServer, featurelabd.NewFeatureLabService(featureLabClient))

	ticker := time.NewTicker(time.Duration(ttlMinutes) * time.Minute)
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err := featureLabClient.UpdateCache()
				if err != nil {
					log.Println("Error updating cache in goroutine: ", err)
				} else {
					log.Println("Automatically updated cache in goroutine")
				}
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}
