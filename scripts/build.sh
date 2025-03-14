#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

cd $DIR/..
go get
go build -o ./pokedexcli-app.exe
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi
echo "Build successful!"