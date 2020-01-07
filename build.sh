#!/bin/bash -x

export GOOS=$1
export GOARCH=amd64

OUTPUTDIR=$2
VERSION=$(git describe --tags)
HASH=$(git rev-parse --short HEAD)
LDFLAGS="-X github.com/chez-shanpu/reposiTree/cmd.Version=${VERSION} -X github.com/chez-shanpu/reposiTree/cmd.Revision=${HASH}"

if [ -z "${GOOS}" ]; then
  echo "GOOS" env is required!
  echo "please run ./build.sh [linux|darwin|windows] [output dir]"
  exit 1
fi

mkdir -p "${OUTPUTDIR}"
go build -ldflags "${LDFLAGS}" -o ./repotr

case "$GOOS" in
"linux" | "darwin")
  SUFFIX=""
  ;;
"windows")
  SUFFIX=".exe"
  ;;
*)
  exit 1
  ;;
esac

mv repotr${SUFFIX} "${OUTPUTDIR}"/repotr-"${GOOS}"-${GOARCH}${SUFFIX}
