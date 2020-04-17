#!/bin/bash
set -e

# get sha hash for unieque container-id
GIT_SHA=$(git rev-parse HEAD)

# google cloud project id (as command line arg)
PROJECT_ID=$1

if [[ -z "$PROJECT_ID" ]]; then
  # exit if PROJECT_ID not provided
  echo "PROJECT_ID not provided"
  exit 1
else
  :
fi

# testing go server
# go test ./...

# compiling
go build -o ./bin/server.out

# building docker container
docker build -t houseofbosons/houseofbosons-blog-v1 .

# pushing to google container registary
docker tag houseofbosons/houseofbosons-blog-v1 gcr.io/$PROJECT_ID/houseofbosons-server-v1:latest
docker tag houseofbosons/houseofbosons-blog-v1 gcr.io/$PROJECT_ID/houseofbosons-server-v1:$GIT_SHA

# pushing to google container registary
docker push gcr.io/$PROJECT_ID/houseofbosons-server-v1:latest
docker push gcr.io/$PROJECT_ID/houseofbosons-server-v1:$GIT_SHA