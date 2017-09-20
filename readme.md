# gRPC Lookaside Load Balancer
This is an implementation of a lookaside (or external/one-arm) load balancer as detailed in the [gRPC documentation for
load balancing](https://grpc.io/blog/loadbalancing). It uses Consul as the primary means of service discovery, but could 
be made to support other applications if needed.

## Usage
Unsurprisingly, this service uses a gRPC-based interface to request addresses (see [proto](_proto/lookaside.proto) for 
services and messages). The application has a simple CLI interface:

```
NAME:
   grpc-lookaside - A lookaside load balancer for gRPC service requests.

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --bind value, -b value        Bind address for the service (default: ":3000")
   --consul value, -c value      Consul address (default: "127.0.0.1:8500")
   --datacenter value, -d value  Consul datacenter (default: "dc1")
   --help, -h                    show help
   --version, -v                 print the version
```