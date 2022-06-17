modules() {
  local module name path git url
  for module in $(@modules); do
    name=$(
      git config --file .gitmodules --get-regexp "submodule\.${module}\.name" |
        awk '{print $2}'
    )
    path=$(
      git config --file .gitmodules --get-regexp "submodule\.${module}\.path" |
        awk '{print $2}'
    )
    git=$(
      git config --file .gitmodules --get-regexp "submodule\.${module}\.url" |
        awk '{print $2}'
    )
    url=$(
      echo "${git}" |
        sed -e 's|:|/|' -e 's|^git@|https://|' -e 's|\.git$||'
    )

    if ! go list -mod=readonly -m | grep -q "^${name}$"; then continue; fi

    local branch tags packages
    branch=$(
      gh repo view --json defaultBranchRef \
        --jq '.defaultBranchRef.name' \
        "${git}"
    )
    tags=$(
      gh repo view --json repositoryTopics \
        --jq '.repositoryTopics? | select(. != null) | map(.name) | join(", ")' \
        "${git}"
    )
    packages=$'\n'
    packages+=$(go list -mod=readonly "./${path}/..." | sed 's|^|    - |')

    cat <<EOF
- prefix: ${name}
  import:
    - url: ${url}
      vcs: git
      source:
        url:  ${url}
        dir:  ${url}/tree/${branch}{/dir}
        file: ${url}/blob/${branch}{/dir}/{file}#L{line}
  packages: ${packages}
  tags: [ ${tags} ]
EOF
  done
}

@modules() {
  git config --file .gitmodules --get-regexp .path |
    awk '{print $2}' |
    sort
}
