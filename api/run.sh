#!/bin/bash

cd ../haoshuju.net
sh run.sh &

cd ../api
go run main.go