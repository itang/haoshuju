#!/bin/bash

cd ../cdn
go run main.go &

cd ../api
go run main.go