#!/bin/sh
set -eu

rm -rf dist/*
mkdir -p dist tmp

go tool templ generate
cp -R public/. dist/
go tool minify -o dist/css/style.css public/css/style.css
go build -o ./tmp/site .
./tmp/site
