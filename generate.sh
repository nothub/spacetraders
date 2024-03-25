#!/usr/bin/env bash

set -o errexit
set -o pipefail

# workdir is repository root
cd "$(dirname "$(realpath "$0")")"

set -o xtrace

if test -n "${1}"; then
    echo >&2 "building tag ${1}"
    ref="tags/${1}"
else
    echo >&2 "building main head"
    ref="heads/main"
fi

set -o nounset

# cleanup
find . | grep -Ev '^\.$|^\./\.git$|^\./\.git/.+|^\./generate.sh$|^\./LICENSE.txt$' | xargs rm -rf

# download docs
tmp="$(mktemp -d)"
curl -fsSL -o "${tmp}/api-docs.zip" "https://github.com/SpaceTradersAPI/api-docs/archive/refs/${ref}.tar.gz"
mkdir -p "${tmp}/api-docs"
tar xvf "${tmp}/api-docs.zip" -C "${tmp}/api-docs" --strip-components=1

# extract version
api_ver="$(cat "${tmp}/api-docs/redocly.yaml" | yq -r '.apis | keys[0]' | sed -n -e 's/^.*@//p')"
api_ref="$(git ls-remote https://github.com/SpaceTradersAPI/api-docs.git | grep "refs/heads/main" | cut -b 1-7)"
pkg_ver="${api_ver}+${api_ref}"

# https://openapi-generator.tech/docs/generators/go/#config-options
gencfg=(
    "packageName=st"
    "packageVersion=${pkg_ver}"

    "isGoSubmodule=false"
    "withGoMod=true"

    "enumClassPrefix=true"
    "structPrefix=true"
    "generateInterfaces=true"
    "generateMarshalJSON=true"
    "hideGenerationTimestamp=true"
    "prependFormOrBodyParameters=false"

    "disallowAdditionalPropertiesIfNotPresent=false"
    "useOneOfDiscriminatorLookup=false"

    "withAWSV4Signature=false"
    "withXml=false"
)

docker run --user "$(id -u):$(id -g)" --rm \
    -v "${PWD}/:/out" \
    -v "${tmp}/api-docs:/docs:ro" \
    openapitools/openapi-generator-cli generate \
        -i "/docs/reference/SpaceTraders.json" \
        -g "go" \
        --git-host "github.com" \
        --git-user-id "nothub" \
        --git-repo-id "spacetraders" \
        --additional-properties "$(
            # shellcheck disable=SC2001
            echo "${gencfg[@]}" | sed 's/ /,/g'
        )" \
        -o "/out"

rm -rf "${tmp}"

go mod tidy
go fmt ./...
go test -v -vet "all" ./...

git add .
git commit -m "Regenerated (spec version ${pkg_ver})"
git push
git tag "v${pkg_ver}"
git push origin "refs/tags/v${pkg_ver}"
