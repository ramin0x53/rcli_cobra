FROM golang:1.18

COPY . /go/src

RUN apt update && apt install -y redis-server

WORKDIR /go/src

RUN make install

ENTRYPOINT redis-server --protected-mode no