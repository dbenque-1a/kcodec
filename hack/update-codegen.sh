#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "--- Generating Groups"
./vendor/k8s.io/code-generator/generate-groups.sh \
    all \
    github.com/dbenque/kcodec/pkg/client \
    github.com/dbenque/kcodec/pkg/api \
    "kcodec:v1 kcodec:v1ext kcodec:v2" \
    --go-header-file ./hack/custom-boilerplate.go.txt

echo "--- Generating Internal Groups"
./vendor/k8s.io/code-generator/generate-internal-groups.sh \
    all \
    github.com/dbenque/kcodec/pkg/client \
    github.com/dbenque/kcodec/pkg/api \
    github.com/dbenque/kcodec/pkg/api \
    "kcodec:v2" \
    --go-header-file ./hack/custom-boilerplate.go.txt
