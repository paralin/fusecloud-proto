#!/bin/bash

migrate -url cassandra://127.0.0.1:9042/global?protocol=4 -path ./ up
