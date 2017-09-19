package lookaside

import (
	"log"
	"strconv"

	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"

	pb "stash.mgmt.local/arch/grpc-lookaside/_proto"
)

type Server struct {
	ConsulAddress    string
	ConsulDatacenter string
	routers          map[string]*Router
}

func (s *Server) Resolve(ctx context.Context, input *pb.Request) (*pb.Response, error) {

	addresses := []string{}
	client, err := api.NewClient(&api.Config{Address: s.ConsulAddress, Datacenter: s.ConsulDatacenter})
	if err != nil {
		log.Printf("error creating client: %v\n", err)
		return nil, err
	}

	services, _, err := client.Catalog().Service(input.Service, "", nil)
	if err != nil {
		log.Printf("error getting catalog services: %v\n", err)
		return nil, err
	}
	var lastService string
	for _, service := range services {
		srvc := service.Address + ":" + strconv.Itoa(service.ServicePort)
		if srvc != lastService {
			addresses = append(addresses, srvc)
			lastService = srvc
		}
	}

	if _, ok := s.routers[input.Service]; !ok {
		s.routers[input.Service] = &Router{Addresses: addresses}
	}

	var response *pb.Response
	switch input.Router {
	case pb.Request_RANDOM:
		response = &pb.Response{Address: s.routers[input.Service].ResolveRandom()}
	case pb.Request_ROUNDROBIN:
		response = &pb.Response{Address: s.routers[input.Service].ResolveRoundRobin()}
	}

	return response, nil
}

func NewServer(address, datacenter string) *Server {
	return &Server{
		ConsulAddress:    address,
		ConsulDatacenter: datacenter,
		routers:          map[string]*Router{},
	}
}
