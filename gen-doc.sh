#!/bin/bash

SRC=$(pwd)
GUNK_SRC=$SRC/gunk

cp base.swagger.json swagger.json

# compile all swaggers path and definitions
for f in $GUNK_SRC/v1/**/*.swagger.json; do
	jq '{paths:.paths,definitions:.definitions}' $f | jq -s '.[0]*.[1]' swagger.json /dev/stdin >s.json && mv s.json swagger.json
done

# generate static html
redoc-cli bundle swagger.json --output=redoc-static-raw.html

rm swagger.json
