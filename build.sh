#! /usr/bin/env bash

function scan() {
  set -ex
  go build -o bin/goscan cmd/scan/main.go
}

function all() {
  set -ex
  go build -o bin/goscan cmd/scan/main.go
}

eval "$@"