#!/bin/bash

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker build -t attack-provider-image .
docker run -p 80:8080 -t attack-provider-image . -d
