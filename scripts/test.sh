#!/usr/bin/env bash

DIR=$(dirname "$0")/..

pushd $DIR
  go get github.com/onsi/ginkgo/ginkgo
  go get github.com/Masterminds/glide
  glide install
  ginkgo -r
popd
