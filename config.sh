#!/bin/bash

# Create a prompt showing the git branch and git dirty status
function prompt_func() {
    local branch=$(git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/')
    if [ ! -z "$branch" ]; then
        local status=$(git status --porcelain 2> /dev/null)
        if [ ! -z "$status" ]; then
            echo -e "\033[0;31m$branch (dirty)\033[0m"
        else
            echo -e "\033[0;32m$branch\033[0m"
        fi
    fi
}

export PS1="\[\033[0;32m\]\u@\h\[\033[0;34m\] \w \$(prompt_func)\[\033[0m\]\$ "

# Add a git alias for a pretty log command showing the commit graph
git config --global alias.lg "log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"

