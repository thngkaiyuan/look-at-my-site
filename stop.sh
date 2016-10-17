#!/usr/bin/env bash

PID_FILE="./server.pid"

if [ -f $PID_FILE ]; then
    PID=$(cat $PID_FILE)
    /bin/kill -9 $PID
    /bin/rm $PID_FILE
fi
