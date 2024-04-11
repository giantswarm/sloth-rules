#!/usr/bin/env bash

# Example how to run it:
# scripts/find-alerts.sh '.labels.team=="atlas"'
# => will report all alerts for team Atlas.

# /!\ This script is provided as-is.
# It won't break anything in your files, but parameters management, help, error handling is missing.
# Meaning: no guarantee about the quality of generated output

# In this place we can file helm-generated rules
rulesFilesDir=helm/sloth-rules/templates
# => prerequisite: have files generated. for instance "make generate" starts with generating files.

# Custom (user-provided) filters
selectQueries=("$@")

# Build `jq` query from filters given as parameters
selectQueriesString="$(printf "| select(%s)\n" "${selectQueries[@]}")"

# For each rules file
for rulesFile in "$rulesFilesDir"/*.yaml; do
    # Retrieve (in an array) alert names that match the query
    mapfile -t alertsList < <(
        yq -ojson "$rulesFile" 2>/dev/null \
        | jq '.spec.slos[].alerting
            '"$selectQueriesString"'
            | .name' 2>/dev/null
    ) || continue

    # Console output
    for alert in "${alertsList[@]}"; do
        echo "$alert"
    done
done
