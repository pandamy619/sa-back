#!/bin/bash

# shellcheck disable=SC2093
# shellcheck disable=SC2046
# shellcheck disable=SC2094
go test $(go list ./... | grep -v /vendor/)

returncode=$?
if [ $returncode -ne 0 ]; then
  echo "tests failed"
  exit 1
fi

