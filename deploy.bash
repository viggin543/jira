#!/usr/bin/env bash
echo "uploading binary to s3"

find . -name "*_test.go" -exec go test {} \;
[[ $? == 1 ]] && exit 1
go build
[[ $? == 1 ]] && echo "build failed" && exit 1
checksum=$(shasum -a 256 jira | awk '{print $1}')
oldChecksum=`grep sha256 ~/personal_projects/opentikva/Jira.rb | grep -o '".*"' | sed s/\"//g`
echo "new version $checksum"
echo "old version $oldChecksum"
[[ $oldChecksum == $checksum ]] && exit 1
current=`pwd`
cd ~/personal_projects/opentikva/
sed -i .bak "s/sha256.*/sha256 \"$checksum\"/"  ./Jira.rb
rm *.bak
git commit -am "bumping jira cli version from $oldChecksum to $checksum" && git push origin master
cd $current
aws s3 --profile personal_s3 cp jira s3://opentikva/ --acl public-read
brew update
brew reinstall jira