#!/bin/bash

cd writstatic/
make
cd ..
docker-compose up -d