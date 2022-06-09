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

# Diffs is a function that shows a git diff between the merge-base and HEAD.
function diffs() {
    local merge_base=$(git merge-base HEAD origin/main)
    git diff $merge_base..HEAD
}

export PS1="\[\033[0;32m\]\u@\h\[\033[0;34m\] \w \$(prompt_func)\[\033[0m\]\$ "
