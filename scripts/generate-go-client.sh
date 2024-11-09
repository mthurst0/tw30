#!/usr/bin/env bash

set -euo pipefail

# ./generate-go-client.sh <input-api-spec-filename> <git-user-id> <git-repo-id> <package-name>

INPUT_API_SPEC=$1
GIT_USER_ID=$2
GIT_REPO_ID=$3
PACKAGE_NAME=$4
PACKAGE_PATH=$5

function usage() {
    echo "$1"
    echo "usage: generate-go-client.sh <input-api-spec-filename> <git-user-id> <git-repo-id> <package-name> <package-path>"
    exit 255
}

if [ ! -f "$INPUT_API_SPEC" ]; then
    usage "file not found: $INPUT_API_SPEC"
fi

if [ -z "$GIT_USER_ID" ]; then
    usage "git user id is required"
fi

if [ -z "$GIT_REPO_ID" ]; then
    usage "git repo id is required"
fi

if [ -z "$PACKAGE_NAME" ]; then
    usage "package name is required"
fi

if [ -z "$PACKAGE_PATH" ]; then
    usage "package path is required"
fi

rm -rfv ./tmp
mkdir -p ./tmp

docker run \
  --user "$(id -u):$(id -g)" \
  --rm \
  --volume "$(pwd):/local" \
  openapitools/openapi-generator-cli:v7.7.0 generate \
  --input-spec "/local/$INPUT_API_SPEC" \
  --generator-name go \
  --template-dir /local/api/openapi-go-client-templates \
  --config /local/api/openapi.generator.config.json \
  --output /local/tmp \
  --git-user-id="${GIT_USER_ID}" \
  --git-repo-id="${GIT_REPO_ID}" \
  --package-name="${PACKAGE_NAME}"

find "${PACKAGE_PATH}" -type f -name "*.go" -delete
cp -rfv ./tmp/*.go "${PACKAGE_PATH}"
rm -rfv ./tmp
