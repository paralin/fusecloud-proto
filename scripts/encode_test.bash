#!/bin/bash

cd ./src/
cat ../data/paralin.doc | protoc --encode auth.User ./auth.proto | od -An -tx1 | tr -d " \n" ; echo
