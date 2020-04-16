#!/bin/sh

go build -o bin/logan-server ./backend/src/main.go 
LOGAN_NATS_HOST=10.211.55.11:4242 LOGAN_MONGO_HOST=10.211.55.11:27017 LOGAN_STORE_IN_DB=true bin/logan-server