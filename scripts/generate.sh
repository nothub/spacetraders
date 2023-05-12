#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

cd "$(dirname "$0")/.."

git submodule update --remote --merge -- "api-docs"
git add "api-docs"
git commit -m "update api specs" -- "api-docs" || true

mkdir -p "api-client"
rm -rf "api-client/{*,.*}"

docker run --user "$(id -u):$(id -g)" --rm                                                 \
    -v "${PWD}/api-client:/api-client"                                                     \
    -v "${PWD}/api-docs:/api-docs:ro"                                                      \
    openapitools/openapi-generator-cli generate                                            \
        -i "/api-docs/reference/SpaceTraders.json"                                         \
        -g "go"                                                                            \
        --git-host "github.com"                                                            \
        --git-user-id "nothub"                                                             \
        --git-repo-id "spacetraders"                                                       \
        --additional-properties "isGoSubmodule=true,packageName=api,enumClassPrefix=true"  \
        -o "/api-client"

find "api-client" -type f ! \( -name '*.go' -o -name 'go.mod' -o -name 'go.sum' -o -name '*.md' \) -delete
find "api-client" -type d -empty -delete

( cd api-client && go mod tidy && go fmt )

git add "api-client"
git commit -m 'update api client' -- "api-client" || true
