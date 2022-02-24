#!/usr/bin/env bash
# shellcheck source=modules.bash # modules

compile() {
  modules | tee modules.yml
  rm -rf dist/*

  pushd dist || exit 1
  maintainer go vanity build -f ../modules.yml
  tree .
  popd || exit 1
}
