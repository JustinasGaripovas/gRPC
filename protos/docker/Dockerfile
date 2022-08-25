FROM golang:1.18-bullseye

RUN apt-get update
RUN apt-get install protobuf-compiler -y
RUN apt-get install make git -y

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2