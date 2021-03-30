# grpcApply

## Install

***grpc***

     go get -u google.golang.org/grpc   

## Build Proto

    protoc --go_out=plugins=grpc:. protos/*.proto
