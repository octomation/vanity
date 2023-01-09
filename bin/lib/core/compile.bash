# shellcheck source=modules.bash # modules

compile() {
  mkdir -p dist && rm -rf dist/*
  modules | tee dist/modules.yml

  pushd dist >/dev/null || exit 1
  maintainer go vanity build -f modules.yml
  tree .
  popd >/dev/null || exit 1
}
