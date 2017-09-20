FROM alpine:3.6
COPY bin/linux_amd64/grpc-lookaside /usr/bin
EXPOSE 3000
CMD /usr/bin/grpc-lookaside -b=$BIND -c=$CONSUL -d=$DATACENTER