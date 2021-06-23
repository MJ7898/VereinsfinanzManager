#!/usr/bin/env bash
export DB_CONNECT=mongodb:27017
docker pull mongo
mkdir -p /mongodata
docker run -it -v mongodata:/data/db -p 27017:27017 --name mongodb -d mongo
docker logs mongodb
