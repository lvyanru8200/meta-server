#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
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

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

GO111MODULE=on go install k8s.io/code-generator/cmd/deepcopy-gen
GO111MODULE=on go install k8s.io/code-generator/cmd/register-gen
GO111MODULE=on go install k8s.io/code-generator/cmd/client-gen
export GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')
export PATH=$PATH:$GOPATH/bin

go_path="${REPO_ROOT}/_go"
cleanup() {
  sudo rm -rf "${go_path}"
}
trap "cleanup" EXIT SIGINT

cleanup

source "${REPO_ROOT}"/hack/util.sh
util:create_gopath_tree "${REPO_ROOT}" "${go_path}"
export GOPATH="${go_path}"

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
echo "Generating with deepcopy-gen"
deepcopy-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/finetune/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/finetune/v1beta1 \
  --output-file-base=zz_generated.deepcopy
deepcopy-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/core/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/core/v1beta1 \
  --output-file-base=zz_generated.deepcopy
deepcopy-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/extension/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/extension/v1beta1 \
  --output-file-base=zz_generated.deepcopy

echo "Generating with register-gen"
register-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/finetune/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/finetune/v1beta1 \
  --output-file-base=zz_generated.register
register-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/core/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/core/v1beta1 \
  --output-file-base=zz_generated.register
register-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-dirs=github.com/DataTunerX/meta-server/apis/extension/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/apis/extension/v1beta1 \
  --output-file-base=zz_generated.register

echo "Generating with client-gen"
client-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-base="" \
  --input=github.com/DataTunerX/meta-server/apis/finetune/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/pkg/generated/clientset \
  --clientset-name=versioned
client-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-base="" \
  --input=github.com/DataTunerX/meta-server/apis/core/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/pkg/generated/clientset \
  --clientset-name=versioned
client-gen \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --input-base="" \
  --input=github.com/DataTunerX/meta-server/apis/extension/v1beta1 \
  --output-package=github.com/DataTunerX/meta-server/pkg/generated/clientset \
  --clientset-name=versioned