#!/usr/bin/env bash

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

DEFAULT_CLUSTER_DOMAIN=example.com
DEFAULT_NAMESPACE=func

show_help() {
  cat << EOF
  Configure a provisioned Kind cluster for use with Functions.

  Requires Kative Serving and Kourier networking.

  Usage: $(basename "$0") <options>

    -h, --help                              Display help
    -n, --namespace                         The namespace to use for Functions (default: $DEFAULT_NAMESPACE)
EOF

}

main() {
  local namespace="$DEFAULT_NAMESPACE"

  parse_command_line "$@"

  namespace 
  network
  kourier_nodeport
  default_domain

  kubectl --namespace kourier-system get service kourier
}

parse_command_line() {
  while :; do
    case "${1:-}" in
      -h|--help)
        show_help
        exit
        ;;
      -n|--namespace)
        if [[ -n "${2:-}" ]]; then
          namespace="$2"
          shift
        else
          echo "ERROR: '-n|--namespace' cannot be empty." >&2
          show_help
          exit 1
        fi
        ;;
      *)
        break
        ;;
    esac
  done
}

namespace() {
  echo 'Creating namespace...'
  kubectl create namespace "$namespace"
  echo 'Marking namespace for eventing injection...'
  kubectl label namespace "$namespace" knative-eventing-injection=enabled
}

network() {
  echo 'Configuring network...'
  kubectl apply -f kind-network.yaml
}

kourier_nodeport() {
  echo 'Configuring nodeport...'
  kubectl patch -n kourier-system services/kourier -p "$(cat kind-kourier-nodeport.yaml)"
}

default_domain() {
  echo 'Configuring default domain...'
  kubectl apply -f kind-domain.yaml
}

main "$@"
