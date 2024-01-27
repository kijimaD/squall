#!/bin/bash
set -eux

#######################
# OpenAPI定義書を結合する
#######################

cd `dirname $0`
cd ../

merger() {
    docker run \
           --rm \
           -v $PWD/oas:/work \
           -w /work \
           redocly/cli:1.4.1 \
           bundle openapi.yml -o merge.yml
}

merger
