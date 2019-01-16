#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'

while true; do
  go build
  go test ./... -timeout=10s
  ./urserver
  #$@ &
  PID=$!
  inotifywait -r -e modify --exclude '\.log' .
  kill $PID
  echo -e "${GREEN}Changes detected! Reloading...${NC}"
done
