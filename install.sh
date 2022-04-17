#!/usr/bin/env bash

export version=0.0.2

if [ -n "$1" ]; then
  version=$1
fi

set -x

echo "> installing goapprove@$version"
curl -LsO https://github.com/amazingandyyy/goapprove/archive/refs/tags/$version.zip &&
unzip -o $version.zip &&
rm -rf /opt/homebrew/bin/goapprove &&
sudo touch /opt/homebrew/bin/goapprove &&
mv -f goapprove-$version/bin/goapprove /opt/homebrew/bin
# rm -rf goapprove-$version goapprove-$version.zip

if ! [ -x "$(command -v goapprove)" ]; then
  echo 'Error: goapprove failed to install' >&2
  exit 1
else
  echo "> install goapprove@$version successfully!"
  goapprove -help
fi

set +x
