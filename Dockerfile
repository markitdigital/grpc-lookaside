FROM alpine:3.6
COPY bin/linux_amd64/grpc-lookaside /usr/bin
COPY docker.sh /app/docker.sh
EXPOSE 3000
CMD [ "sh", "/app/docker.sh" ]