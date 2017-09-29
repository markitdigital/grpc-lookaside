package lookaside

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"

	pb "github.com/markitondemand/grpc-lookaside/_proto"
)

type Server struct {
	ConsulAddress    string
	ConsulDatacenter string
	routers          map[string]*Router
	refreshInterval  float64
}

func (s *Server) Resolve(ctx context.Context, input *pb.Request) (*pb.Response, error) {

	var (
		address string
		err     error
	)

	// ensure router exists and addresses are fresh
	if router, ok := s.routers[input.Service]; !ok || router.NeedsRefresh() {
		addresses, err := s.refreshAddresses(input.Service)
		if err != nil {
			return nil, err
		}

		s.routers[input.Service] = &Router{Addresses: addresses, LastRefresh: time.Now(), RefreshInterval: s.refreshInterval}
	}

	// determine the type of routing requested, and resolve an address
	switch input.Router {
	case pb.Request_RANDOM:
		address, err = s.routers[input.Service].ResolveRandom()
	case pb.Request_ROUNDROBIN:
		address, err = s.routers[input.Service].ResolveRoundRobin()
	case pb.Request_HASH:
		address, err = s.routers[input.Service].ResolveHash(input.Hash)
	}

	if err != nil {
		return &pb.Response{Address: ""}, err
	}

	return &pb.Response{Address: address}, nil
}

func (s *Server) refreshAddresses(service string) ([]string, error) {

	// create a "set" using a map[string]struct{} to hold unique addresses
	addressSet := make(map[string]struct{})

	// create a consul client
	consul, err := api.NewClient(&api.Config{Address: s.ConsulAddress, Datacenter: s.ConsulDatacenter})
	if err != nil {
		return make([]string, 0), err
	}

	// retrieve service catalog members
	members, _, err := consul.Catalog().Service(service, "", nil)
	if err != nil {
		return make([]string, 0), err
	}

	// loop through members and build up set of addresses
	for _, member := range members {
		addressSet[fmt.Sprintf("%s:%s", member.Address, strconv.Itoa(member.ServicePort))] = struct{}{}
	}

	// create slice from set map keys
	addresses := make([]string, 0, len(addressSet))
	for k := range addressSet {
		addresses = append(addresses, k)
	}

	return addresses, nil

}

func NewServer(address, datacenter string, refresh float64) *Server {
	return &Server{
		ConsulAddress:    address,
		ConsulDatacenter: datacenter,
		routers:          map[string]*Router{},
		refreshInterval:  refresh,
	}
}
