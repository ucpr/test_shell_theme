#!/usr/bin/zsh
#
# usage
# -----
# If you are developing, replace "HINA_PATH" with your PATH. 
#
# $ source develop.sh
#

export HINA_PATH="/home/ucpr/Works/hina/hina"

autoload -Uz add-zsh-hook
_hina_prompt() {
  PROMPT=$($HINA_PATH)
}

add-zsh-hook precmd _hina_prompt
