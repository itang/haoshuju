#!/bin/bash

if [ "$1" = "" ]; then
  bee run haoshuju.com
else
  RunMode="$1" bee run haoshuju.com
fi
