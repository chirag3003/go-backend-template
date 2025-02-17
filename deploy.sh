#!/bin/sh

docker stop GoBackend;
docker rm  GoBackend;

git pull;

docker build . -t go-backend;
docker run --name GoBackend -p 5000:5000 -d go-backend;
