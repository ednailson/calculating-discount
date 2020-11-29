#!/usr/bin/env bash

DEFAULT_LIMIT=60
waitContainerUp()
{
    SERVICE=$1
    LIMIT=$2
    NEXT_WAIT_TIME=0
    until docker container inspect --format="${SERVICE} is {{.State.Status}}" $(docker-compose ps -q ${SERVICE} || echo "0") | grep -i "running" || [ $NEXT_WAIT_TIME -eq $LIMIT ]; do
       sleep 1
       echo "retry wait ${SERVICE} up: $(( ++NEXT_WAIT_TIME ))"
    done
}

docker-compose up -d --build
waitContainerUp arango $DEFAULT_LIMIT
docker-compose exec arango sh opt/tools/db.sh

docker-compose logs -f --tail="100"