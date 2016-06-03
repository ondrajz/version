#!/bin/sh
set -e

VERSION_PKG=${1:-`go list`/version}
BASE_VERSION=$(git describe --always --dirty=-dev | sed 's/-\([0-9]*\)-g[0-9a-f]*/+\1/')
COMMIT_HASH=$(git show -s --format=%h)
COMMIT_STAMP=$(git show -s --format=%ct)
BUILD_USER=$(id -u -n)
BUILD_HOST=$(hostname)
BUILD_STAMP=$(date +%s)

vflag() {
    VFLAGS="$VFLAGS -X $VERSION_PKG.$1=$2"
}

vflag Number $BASE_VERSION
vflag CommitHash $COMMIT_HASH
vflag CommitStamp $COMMIT_STAMP
vflag BuildUser $BUILD_USER
vflag BuildHost $BUILD_HOST
vflag BuildStamp $BUILD_STAMP

echo "$VFLAGS"
