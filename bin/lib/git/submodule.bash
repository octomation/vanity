# issue: https://github.com/kamilsk/dotfiles/issues/601
@git-submodule-rm() {
  local submodule="${1}"

  git submodule status "${submodule}" || return 1
  git --no-pager config -f .gitmodules --get-regexp "submodule.${submodule}"

  git submodule deinit -f "${submodule}"
  rm -rf ".git/modules/${submodule}"
  git rm -f "${submodule}"
  git config -f .gitmodules --remove-section "submodule.${submodule}"
}
