#!/bin/bash

docker-compose build
docker-compose up -d
docker-compose exec boiler-plate ash
