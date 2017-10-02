package main

import (
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/markitondemand/grpc-lookaside"
	pb "github.com/markitondemand/grpc-lookaside/_proto"
)

func main() {

	app := cli.NewApp()
	app.Name = "grpc-lookaside"
	app.Usage = "A lookaside load balancer for gRPC service requests."
	app.Version = "0.0.6"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind,b",
			Value: ":3000",
			Usage: "Bind address for the service",
		},
		cli.StringFlag{
			Name:   "consul,c",
			Value:  "127.0.0.1:8500",
			Usage:  "The Consul address used when querying for services, typically an agent on the same machine",
			EnvVar: "CONSUL_ADDRESS",
		},
		cli.StringFlag{
			Name:   "datacenter,d",
			Value:  "dc1",
			Usage:  "The Consul datacenter used when querying for services",
			EnvVar: "CONSUL_DATACENTER",
		},
		cli.Float64Flag{
			Name:   "refresh,r",
			Value:  10.00,
			Usage:  "Time, in seconds, between service address list refreshes",
			EnvVar: "REFRESH",
		},
	}
	app.Action = func(c *cli.Context) error {

		// create a TCP listener on the provided bind address
		listener, err := net.Listen("tcp", c.String("bind"))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Printf("Listening for requests at %s\n", c.String("bind"))
		defer listener.Close()

		// register the server implementation with the generated handlers and listen for incoming requests
		server := grpc.NewServer()
		pb.RegisterLookasideServer(server, lookaside.NewServer(c))
		return server.Serve(listener)

	}

	log.Fatal(app.Run(os.Args))
}
