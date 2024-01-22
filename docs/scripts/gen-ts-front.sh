#!/bin/bash
set -eux

##################################
# OpenAPIからフロントのコードを生成する
##################################

cd `dirname $0`

mkdir -p ../../front/src/renderer/generated

docker run --rm \
       -v $PWD/../..:/work \
       -w /work/docs/scripts \
       openapitools/openapi-generator-cli:latest-release generate \
         -i /work/docs/oas/openapi.yml \
         -g typescript-axios \
         -o /work/front/src/renderer/generated
