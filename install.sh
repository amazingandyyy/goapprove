#!/usr/bin/env bash

export version=0.0.1
curl -LsO https://github.com/amazingandyyy/goapprove/archive/refs/tags/$version.zip &&
unzip -o $version.zip &&
sudo touch /opt/homebrew/bin/goapprove &&
mv -f goapprove-$version/bin/goapprove /opt/homebrew/bin &&
rm -rf goapprove-$version $version.zip

echo install goapprove@$version success! try 'goapprove --url'
