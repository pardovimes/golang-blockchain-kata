#!/bin/bash

docker run --rm -ti -w="/go/src" \
    -v $PWD/docker/.composer:/home/wwwagent/.composer:delegated \
    -v $PWD:/go/src/blockgo golang /bin/bash
