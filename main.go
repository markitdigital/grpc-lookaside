package main

import (
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	pb "stash.mgmt.local/arch/grpc-lookaside/_proto"
	"stash.mgmt.local/arch/grpc-lookaside/lib"
)

func main() {

	app := cli.NewApp()
	app.Name = "grpc-lookaside"
	app.Usage = "A lookaside load balancer for gRPC service requests."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind,b",
			Value: ":3000",
			Usage: "Bind address for the service",
		},
		cli.StringFlag{
			Name:  "consul,c",
			Value: "127.0.0.1:8500",
			Usage: "Consul address",
		},
		cli.StringFlag{
			Name:  "datacenter,d",
			Value: "dc1",
			Usage: "Consul datacenter",
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
		pb.RegisterLookasideServer(server, lookaside.NewServer(c.String("consul"), c.String("datacenter")))
		return server.Serve(listener)

	}

	log.Fatal(app.Run(os.Args))
}
