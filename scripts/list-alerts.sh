#!/usr/bin/env bash

# Example how to run it:
# scripts/list-alerts.sh '.labels.team=="atlas"'
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
    # As of now, the slos field in the prometheusServiceLebel Resources only contains one alert because of the way those are generated with our abstraction layer.
    # If the way we write slos was to change and multiple alerts could be created in the slos field, this would have to be updated accordingly.
    alertList="$(yq -ojson "$rulesFile" 2>/dev/null | jq '.spec.slos[].alerting'"$selectQueriesString"' | .' 2>/dev/null)"
    
    if [[ $alertList != "" ]]; then
        name="$(echo $alertList | jq -r '.name' 2>/dev/null)"
        team="$(echo $alertList | jq -r '.labels.team?' 2>/dev/null)"
        isSilenced="$(echo $alertList | jq -r '.pageAlert.labels.silence?' 2>/dev/null)"

        if [[ $isSilenced != "true" ]]; then
            isSilenced="false"
        fi
        
        echo "alert name : \"$name\" - team : \"$team\" - silence : \"$isSilenced\""
    fi 

done
