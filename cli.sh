#!/bin/bash

docker run --rm -ti -w="/app" \
    -v $PWD/docker/.composer:/home/wwwagent/.composer:delegated \
    -v $PWD:/app golang /bin/bash
