#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

cd accountservice;go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..
cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..

export GOOS=windows

cp healthchecker/healthchecker-linux-amd64 accountservice/

docker build -t goblog/accountservice accountservice/

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=goblog -p=3000:3000 goblog/accountservice