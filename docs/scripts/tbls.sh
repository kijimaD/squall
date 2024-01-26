#!/bin/bash
set -eux

############################################
# tblsでデータベースドキュメントを生成する
############################################

cd `dirname $0`
cd ../..

lint_tbls() {
    docker run \
           --env TBLS_DSN=sqlite:///work/api/squall.sqlite \
           --env TBLS_DOC_PATH=docs/dbschema \
           --rm \
           -v $PWD:/work \
           -w /work \
           ghcr.io/k1low/tbls \
           lint \
           -c docs/.tbls.yml
}

lint_tbls

generate_tbls() {
    docker run \
           --env TBLS_DSN=sqlite:///work/api/squall.sqlite \
           --env TBLS_DOC_PATH=docs/dbschema \
           --rm \
           -v $PWD:/work \
           -w /work \
           ghcr.io/k1low/tbls \
           doc \
           -c docs/.tbls.yml \
           --force
}

generate_tbls
