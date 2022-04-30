@warmup() {
  local module=${1} version=${2}
  curl -v "https://proxy.golang.org/${module}/@v/${version}.info"
}

@test() {
  local module=${1} version=${2} tmp

  tmp=$(mktemp -d)
  pushd "${tmp}" >/dev/null || exit 1
  go mod init test
  GOPROXY=https://proxy.golang.org GO111MODULE=on go get "${module}@${version}"
  popd >/dev/null || exit 1
}
