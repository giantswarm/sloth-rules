#!/bin/bash

main() {
  local GIT_WORKDIR
  GIT_WORKDIR="$(git rev-parse --show-toplevel)"

  echo "Templating chart"

  helm template \
    "$GIT_WORKDIR"/helm/sloth-rules \
    --output-dir "$GIT_WORKDIR"/scripts/tests/output/
}

main "$@"
