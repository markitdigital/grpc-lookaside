package lookaside

import (
	"testing"
)

func TestResolveHash(t *testing.T) {

	router := Router{Addresses: []string{"127.0.0.1:3000", "127.0.0.1:3001"}}

	var tests = []struct {
		input    string
		expected string
	}{
		{"foo", "127.0.0.1:3000"}, // first up
		{"foo", "127.0.0.1:3000"}, // repeat for same key
		{"bar", "127.0.0.1:3001"}, // round robin, next address
	}

	for _, test := range tests {
		addr, _ := router.ResolveHash([]byte(test.input))
		if addr != test.expected {
			t.Errorf("Router.ResolveHash(%v): expected %v, actual %v", test.input, test.expected, addr)
		}
	}

}

func TestResolveRandom(t *testing.T) {

	addresses := make(map[string]struct{})
	router := Router{Addresses: []string{"127.0.0.1:3000", "127.0.0.1:3001", "127.0.0.1:3002"}}
	for i := 0; i < 100; i++ {
		addr, _ := router.ResolveRandom()
		addresses[addr] = struct{}{}
	}

	if len(addresses) != 3 {
		t.Errorf("Router.ResolveRandom(): Expected 3 unique addresses over 100 requests, got %v. Unlikely, but could be a red herring.", len(addresses))
	}
}

func TestResolveRoundRobin(t *testing.T) {

	router := Router{Addresses: []string{"127.0.0.1:3000", "127.0.0.1:3001", "127.0.0.1:3002"}}

	var tests = []string{"127.0.0.1:3000", "127.0.0.1:3001", "127.0.0.1:3002", "127.0.0.1:3000"}

	for _, test := range tests {

		addr, _ := router.ResolveRoundRobin()
		if addr != test {
			t.Errorf("Router.ResolveRoundRobin(): expected %v, actual %v", test, addr)
		}

	}

}
