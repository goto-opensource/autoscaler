#!/bin/bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

echo "verify-kubelint"

cd "$(dirname "${BASH_SOURCE[0]}")/.."
SCRIPT_ROOT="$PWD"

TOOLS_DIR="${PWD}/hack/tools/kube-api-linter"
TOOLS_BIN_DIR="${TOOLS_DIR}/bin"

# Use the custom builder's pinned version, not a newer CLI requiring a newer Go.
GOLANGCI_LINT_VERSION=$(awk '$1 == "version:" { print $2; exit }' "${TOOLS_DIR}/.custom-gcl.yml")
if [[ ! "${GOLANGCI_LINT_VERSION}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Missing or invalid golangci-lint version in ${TOOLS_DIR}/.custom-gcl.yml" >&2
  exit 1
fi

echo "installing dependencies"
go install "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}"

GOLANGCI_LINT_BIN=${GOLANGCI_LINT_BIN:-"golangci-lint"}
GOLANGCI_LINT_KAL_BIN=${GOLANGCI_LINT_KAL_BIN:-"${TOOLS_BIN_DIR}/golangci-kube-api-linter"}
GOLANGCI_LINT_CONFIG_PATH=${GOLANGCI_LINT_CONFIG_PATH:-"${TOOLS_DIR}/.golangci-kal.yml"}

echo "creating custom golangci linter"
cd "${TOOLS_DIR}"; "${GOLANGCI_LINT_BIN}" custom
cd "${SCRIPT_ROOT}"

"${GOLANGCI_LINT_KAL_BIN}" run -v --config "${GOLANGCI_LINT_CONFIG_PATH}"
