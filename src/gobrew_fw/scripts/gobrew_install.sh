#!/usr/bin/env bash

set -o errtrace
export PS4='+[${BASH_SOURCE}] : ${LINENO} : ${FUNCNAME[0]:+${FUNCNAME[0]}() $ }'

#if [[ -f "$HOME/.rvmrc" ]] ; then source "$HOME/.rvmrc" ; fi

gobrew_path="${gobrew_path:-$HOME/.gobrew}"

mkdir -p $gobrew_path

builtin cd $gobrew_path

stable_version=$(curl -B <%= HOST %>/releases/stable-version.txt 2>/dev/null)

echo "Installing RVM Version: ${stable_version}"

curl -L "<%= HOST %>/rubies/packages/rvm/${stable_version}.tar.gz" -o "${stable_version}.tar.gz"


# set name of folder we cd into after file is extracted
rvm_src_dir=$(tar tzf "${stable_version}.tar.gz" | tail -1 | awk '{split ($0, a, "/"); print a[1]}')
# extract the downloaded tarball
tar zxf "${stable_version}.tar.gz"
# cd into the extracted directory
builtin cd "${rvm_src_dir}"

# INSTALL RVM!!!
bash ./scripts/install
