#!/bin/bash

SRC=$(pwd)
GUNK_SRC=$SRC/gunk

pushd $SRC &>/dev/null

# loop through gunk definition and generate pb, gw and swagger files.
pushd $GUNK_SRC/v1 &> /dev/null

for i in */*.gunk; do
    pushd $(dirname $i) &> /dev/null
    
    f=$(basename $i)
    d=${f%.*}
    gunk generate $f

    # delete empty swagger
    if grep -q -E '"paths": {}|"definitions": {}' all.swagger.json; then
        rm all.swagger.json
    fi

    # rename generated files
    for a in all.*; do
        mv "$a" "$(echo "$a" | sed s/all/$d/)"
    done

    popd &>/dev/null 
done

popd &>/dev/null 
