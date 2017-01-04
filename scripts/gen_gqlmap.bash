#!/bin/bash
set -e

go run ./scripts/gqlreflect.go

find . -name 'graphql.types.go' -exec gofmt -w -s {} \;
