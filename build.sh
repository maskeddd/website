#!/bin/sh
set -eu

mkdir -p dist/css tmp

go tool templ generate
cp -R public/. dist/
find dist -type f -name .gitkeep -delete
tailwindcss -i ./assets/css/input.css -o ./dist/css/style.css
go build -o ./tmp/site .
./tmp/site
