#!/bin/bash

set -euxo pipefail

function build() {
  GOOS="$1"
  GOARCH="$2"
  COMMAND="$3"

  go build -o "$OUTDIR/$GOOS-$GOARCH/$COMMAND" "./cmd/$COMMAND"
}

function package() {
  GOOS="$1"
  GOARCH="$2"

  pushd "$OUTDIR/"
  zip --recurse-paths "./$GOOS-$GOARCH.zip" "./$GOOS-$GOARCH/"
  popd

  rm --recursive "$OUTDIR/$GOOS-$GOARCH/"
}

CGO_ENABLED=0
DEFAULT_OUTDIR="$(dirname ${BASH_SOURCE[0]})/../dist"
OUTDIR="${OUTDIR:-$DEFAULT_OUTDIR}"

build darwin amd64 spaced
build linux amd64 spaced
build windows amd64 spaced

package darwin amd64
package linux amd64
package windows amd64
