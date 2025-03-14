#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"

echo "Current directory: $(pwd)"
echo "Running build.sh..."

$DIR/build.sh && $DIR/../pokedexcli-app.exe 

