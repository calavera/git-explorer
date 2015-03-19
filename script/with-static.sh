#!/bin/sh

set -ex

LIBGIT2="$GOPATH/src/github.com/libgit2/git2go/vendor/libgit2"

export BUILD="$LIBGIT2/build"
export PCFILE="$BUILD/libgit2.pc"

FLAGS=$(pkg-config --static --libs $PCFILE) || exit 1
export CGO_LDFLAGS="$BUILD/libgit2.a -L$BUILD ${FLAGS}"
export CGO_CFLAGS="-I$LIBGIT2/include"

$@
