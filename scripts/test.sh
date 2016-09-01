#!/usr/bin/env bash

set -xe

DIR=$(dirname "$0")/..

echo $GOPATH

pushd $DIR
  go get github.com/onsi/ginkgo/ginkgo
  go get github.com/Masterminds/glide
  glide install
  ginkgo -r
popd
