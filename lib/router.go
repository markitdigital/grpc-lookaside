package lookaside

import (
	"log"
	"math/rand"
	"time"
)

type Router struct {
	Addresses           []string
	roundRobinLastIndex int
}

func (r *Router) ResolveRandom() string {
	rand.Seed(time.Now().Unix())
	if len(r.Addresses) == 0 {
		log.Println("Attempted to resolve address using RandomRouter with empty array.")
		return ""
	}
	return r.Addresses[rand.Intn(len(r.Addresses))]
}

func (r *Router) ResolveRoundRobin() string {
	if r.roundRobinLastIndex == (len(r.Addresses) - 1) {
		r.roundRobinLastIndex = 0
	} else {
		r.roundRobinLastIndex++
	}
	return r.Addresses[r.roundRobinLastIndex]
}
