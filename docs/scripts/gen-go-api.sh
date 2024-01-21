#!/bin/bash
set -eux

#################################
# OpenAPIからGoのAPIコードを生成する
#################################

cd `dirname $0`

mkdir -p ../../api/generated

docker run \
       --rm \
       -v $PWD/../../:/work \
       -w /work/docs/scripts \
       golang:1.21 \
       bash -c "go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
       oapi-codegen --config ../oas/config/server.yml ../oas/merge.yml && \
       oapi-codegen --config ../oas/config/models.yml ../oas/merge.yml && \
       oapi-codegen --config ../oas/config/spec.yml ../oas/merge.yml"
