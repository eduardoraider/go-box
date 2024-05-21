#!/bin/bash

apps=("api:api" "gobox:cli" "worker:worker")

for app in "${apps[@]}"; do
  artefact="${app%%:*}"
  app="${app#*:}"

  echo "building $artefact"
  go build -o "${artefact}" ./cmd/"${app}"
done
