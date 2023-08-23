#!/usr/bin/env bash
cat << _EOF_ | nc localhost 8000
GET /hijack HTTP/1.1
Host: localhost:8000

Hi server!
_EOF_