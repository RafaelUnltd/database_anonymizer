FROM golang:alpine

# Adiciona repositório do PostgreSQL 17
RUN echo "https://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories
RUN echo "https://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk update
RUN apk add --no-cache ca-certificates git openssh-client && rm -rf /var/cache/apk/*
RUN apk add --no-cache postgresql17-client
RUN mkdir /root/.ssh/
RUN mkdir -p /tmp/database_dumps

WORKDIR app/

RUN go install github.com/githubnemo/CompileDaemon@v1.4.0

EXPOSE 4001

CMD CompileDaemon --directory="." --build="go build app/cmd/main.go" --command=./main --exclude-dir="/tmp/database_dumps" --pattern="(.+\.go|.+\.c|.+\.sql)$"
