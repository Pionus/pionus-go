#! /bin/bash

OUTFILE="pionus"
PIDFILE="server.pid"

if [ -f $PIDFILE ]; then
    kill -9 $(cat $PIDFILE)

go build -o $OUTFILE server.go
./$OUTFILE & echo $! > $PIDFILE
