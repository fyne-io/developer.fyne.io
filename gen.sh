#!/bin/sh

ROOT="`go env GOPATH`/src/fyne.io/fyne"

# generate API docs

DIRS=`find $ROOT -not -path '*/\.*' -type d | grep -v vendor | grep -v internal`
PREFIX="api"

godocdown -template="_gen/api.md" $ROOT > "$PREFIX/index.md" 2>&1 | grep -v "Could not find package"
for DIR in $DIRS; do
  PKG=`echo $DIR | cut -c$((${#ROOT}+2))-`
 
  mkdir -p "$PREFIX/$PKG"
 
  godocdown -template="_gen/api.md" -heading="1Word" $DIR > "$PREFIX/$PKG/index.md" 2>&1 | grep -v "Could not find package"
done
