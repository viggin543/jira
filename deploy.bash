#!/usr/bin/env bash
echo "uploading binary to s3"

find . -name "*_test" | xargs -I % go test % && \
go build && \
aws s3 --profile personal_s3 cp jira s3://opentikva/ --acl public-read