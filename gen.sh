#!/bin/sh

go get -u github.com/andydotxyz/godocdown/godocdown

ROOT="`go env GOPATH`/src/fyne.io/fyne"

# generate API docs

DIRS=`find $ROOT -not -path '*/\.*' -type d | grep -v vendor | grep -v internal | grep -v testdata | grep -v cmd`
PREFIX="api"

godocdown -template="_gen/api.md" -outputDir "$PREFIX/" $ROOT 2>&1 | grep -v "Could not find package"
for DIR in $DIRS; do
  PKG=`echo $DIR | cut -c$((${#ROOT}+2))-`
 
  mkdir -p "$PREFIX/$PKG"
 
  godocdown -template="_gen/api.md" -outputDir "$PREFIX/$PKG/" $DIR 2>&1 | grep -v "Could not find package"
done
