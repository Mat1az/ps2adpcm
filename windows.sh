#!/bin/bash

statik -f src=public/win_x86_64
./common.sh
env GOOS=windows GOARCH=amd64 go build ps2adpcm