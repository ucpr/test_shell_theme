#!/usr/bin/zsh
#
# usage
# -----
# If you are developing, replace "HINA_PATH" with your PATH. 
#
# $ source develop.sh
#

export HINA_GIT_MODIFIED="*"
export HINA_GIT_ADDED="+"
export HINA_GIT_DELETED="-"
export HINA_GIT_RENAMED="$"
export HINA_GIT_COPIED="~"
export HINA_GIT_UNMERGED="="
export HINA_GIT_UNTRACKED="!"

export HINA_PATH="/home/ucpr/Works/hina/hina"

autoload -Uz add-zsh-hook
_hina_prompt() {
  PROMPT=$($HINA_PATH)
}

add-zsh-hook precmd _hina_prompt
