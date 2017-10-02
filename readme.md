# gRPC Lookaside Load Balancer
This is an implementation of a lookaside (or external/one-arm) load balancer as detailed in the [gRPC documentation for
load balancing](https://grpc.io/blog/loadbalancing) that uses Consul for service discovery. It supports multiple types of load-balancing (random, round-robin, hash) and periodic address refreshes.

## Usage
Unsurprisingly, this service uses a gRPC-based interface to request addresses (see [proto](_proto/lookaside.proto) for 
services and messages). 

## CLI
The application has a simple CLI interface and supports the following arguments and environment variable configurations:

### --bind, b
The address that the service will bind to in {host}:{port} format. Defaults to `:3000`.

### --consul, c [$CONSUL_ADDRESS]
The address of the Consul agent used for service discovery. Defaults to `127.0.0.1:8500` (local agent).

### --datacenter, d [$CONSUL_DATACENTER]
The Consul datacenter to query for services. Defaults to `dc1`.

### --refresh, r [$REFRESH]
Time, in seconds, between service address list refreshes. Defaults to 10 seconds.

## Building
You'll need a working install of the Go programming language to compile the go code and the GNU `make` tool for running
the builds. Once everything is installed and setup, just run:

```bash
$ make
```