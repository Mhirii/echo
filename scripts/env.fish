#!/usr/bin/env fish

# set env_path $argv[1]
set env_path .env

cat $env_path | sed 's/^\([^=]*\)=\(.*\)$/set \1 "\2"/' > .env.fish
