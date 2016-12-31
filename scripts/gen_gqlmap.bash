#!/bin/bash
set -e

go run ./scripts/gqlreflect.go
gofmt -w -s root_gql.go
