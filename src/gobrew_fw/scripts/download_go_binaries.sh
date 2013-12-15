#!/usr/bin/env bash

version=1.2
downloads_url=https://go.googlecode.com/files
local_dir=web/files

# mac osx 10.6
wget -P "$local_dir"  "$downloads_url/go$version.darwin-386-osx10.6.tar.gz"
wget -P "$local_dir"  "$downloads_url/go$version.darwin-amd64-osx10.6.tar.gz"

# mac osx 10.8
wget -P "$local_dir"  "$downloads_url/go$version.darwin-386-osx10.8.tar.gz"
wget -P "$local_dir"  "$downloads_url/go$version.darwin-amd64-osx10.8.tar.gz"

# bsd
wget -P "$local_dir"  "$downloads_url/go$version.freebsd-386.tar.gz"
wget -P "$local_dir"  "$downloads_url/go$version.freebsd-amd64.tar.gz"

# linux
wget -P "$local_dir"  "$downloads_url/go$version.linux-386.tar.gz"
wget -P "$local_dir"  "$downloads_url/go$version.linux-amd64.tar.gz"

# windows
wget -P "$local_dir"  "$downloads_url/go$version.windows-386.zip"
wget -P "$local_dir"  "$downloads_url/go$version.windows-amd64.zip"
