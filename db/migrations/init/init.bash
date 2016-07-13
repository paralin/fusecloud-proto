#!/bin/bash
CASSANDRA_IP=10.244.2.4
migrate -url cassandra://$CASSANDRA_IP:9042/global?protocol=4 -path ./ up
