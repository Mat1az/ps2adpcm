#!/bin/bash

statik src=public/linux_x86_64
./common.sh
env GOOS=linux GOARCH=amd64 go build