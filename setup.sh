#! /usr/bin/env bash

VIQUEEN_GO_DEVBOX_HOME=$(cd "$(dirname "$0")" && pwd -P)

_shell_rc() {
  if [[ -n $ZSH_VERSION ]]; then
    echo "${HOME}/.zshrc"
  elif [[ -n $FSH_VERSION ]]; then
    echo "${HOME}/.fshrc"
  else # default to bash
    echo "${HOME}/.bashrc"
  fi
}

function binary() {
  set -ex
  local -r rc_file=$(_shell_rc)
  echo "export VIQUEEN_GO_DEVBOX_HOME=${VIQUEEN_GO_DEVBOX_HOME}" >> "${rc_file}"
  echo "export PATH=${PATH}:${VIQUEEN_GO_DEVBOX_HOME}/bin" >> "${rc_file}"
}

eval "$@"
