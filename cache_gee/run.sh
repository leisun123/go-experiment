#!/bin/zsh
trap "rm server;kill 0" EXIT

go build -o server
./server -port=7001 &
./server -port=7002 &
./server -port=7003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" & curl "http://localhost:9999/api?key=Tom" & curl "http://localhost:9999/api?key=Tom"

wait