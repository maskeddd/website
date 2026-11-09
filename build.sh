#!/bin/sh
set -eu

rm -rf dist/*
mkdir -p dist/css tmp

go tool templ generate
cp -R public/. dist/
tailwindcss -i ./assets/css/input.css -o ./dist/css/style.css
go build -o ./tmp/site .
./tmp/site
