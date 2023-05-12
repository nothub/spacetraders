#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

cd "$(dirname "$0")/.."

if test $# -gt 0; then
    symbol="$(echo "${1}" | tr '[:lower:]' '[:upper:]')"
else
    symbol=$(pwgen --ambiguous --num-passwords 1 | tr '[:lower:]' '[:upper:]')
fi

faction="$(curl --silent --show-error --location \
    --request GET \
    --url "https://api.spacetraders.io/v2/factions" \
    --header "Accept: application/json" |
    jq -r '.data[].symbol' | shuf -n 1)"

curl --silent --show-error --location \
    --request POST \
    --url "https://api.spacetraders.io/v2/register" \
    --header "Accept: application/json" \
    --header "Content-Type: application/json" \
    --data "{\"symbol\":\"${symbol}\",\"faction\":\"${faction}\"}" |
    jq >"agent-${symbol}.json"

echo "registered agent ${symbol} for faction ${faction}" >&2
echo "token: $(jq -r '.data.token' "agent-${symbol}.json")" >&2
