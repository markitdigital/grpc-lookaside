package lookaside

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Router is a type that contains a set of addresses and methods for resolving them based on use input
type Router struct {
	Addresses        []string
	LastRefresh      time.Time
	RefreshInterval  float64
	roundRobinCursor int
	hashRouteMap     map[string]string
}

// ResolveHash returns an address bound to the hash provided, or picks a new one and binds it
func (r *Router) ResolveHash(data []byte) (string, error) {

	// initialize the hashRouteMap if needed
	if r.hashRouteMap == nil {
		r.hashRouteMap = make(map[string]string)
	}

	// hash data
	hash := fmt.Sprintf("%x", md5.Sum(data))

	// if the hash is already bound to an address, use that
	if _, ok := r.hashRouteMap[hash]; !ok {

		// resolve a new address for the hash
		address, err := r.ResolveRoundRobin()
		if err != nil {
			return "", err
		}

		// bind the address to the hash
		r.hashRouteMap[hash] = address
	}

	return r.hashRouteMap[hash], nil

}

// ResolveRandom returns a random address
func (r *Router) ResolveRandom() (string, error) {

	if len(r.Addresses) == 0 {
		return "", errors.New("attempted to resolve address with empty array")
	}

	return r.Addresses[rand.Intn(len(r.Addresses))], nil

}

// ResolveRoundRobin returns the next address, looping back to the first when reaching the end
func (r *Router) ResolveRoundRobin() (string, error) {

	if len(r.Addresses) == 0 {
		return "", errors.New("attempted to resolve address with empty array")
	}

	// if the cursor is out of range, reset it to the first item
	if r.roundRobinCursor == len(r.Addresses) {
		r.roundRobinCursor = 0
	}

	// resolve address, increment cursor position, return
	address := r.Addresses[r.roundRobinCursor]
	r.roundRobinCursor++
	return address, nil

}

func (r *Router) NeedsRefresh() bool {
	return time.Now().Sub(r.LastRefresh).Seconds() >= r.RefreshInterval
}
