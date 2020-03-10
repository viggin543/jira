#!/usr/bin/env bash
echo "uploading binary to s3"

go test ./... && go build || ≠(echo "build failed" && exit 1)


mkdir target
mv jira target && cp completions.bash target
tar -zcvf jira_cli.tar.gz target

checksum=$(shasum -a 256 jira_cli.tar.gz | awk '{print $1}')
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

aws s3 --profile personal_s3 cp jira_cli.tar.gz s3://opentikva/ --acl public-read
brew update
brew reinstall jira
rm -rf target
rm -rf jira_cli.tar.gz