#!/usr/bin/env bash

PID_FILE="./server.pid"
LOG_FILE="./server.log"
ERR_FILE="./server.err"

make deps
make build

if [ -f $PID_FILE ]; then
    PID=$(cat $PID_FILE)
    /bin/kill -9 $PID
    /bin/rm $PID_FILE
fi

./look-at-my-site > $LOG_FILE 2> $ERR_FILE &
PID=$!
echo $PID > $PID_FILE
