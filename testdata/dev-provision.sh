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

# *****************
# **** WARNING ****
# This is a work-in-progress, and is neither complete nor entirely safe.
# *****************
#
# Provision a kind cluster with Knative and Kourier installed.
# Suitable for use locally during development.
# CI uses knative-kind action

set -o errexit
set -o nounset
set -o pipefail

main() {

  local serving_version=v0.20.0
  local eventing_version=v0.20.0
  local kourier_version=v0.20.0

  cluster
  serving
  eventing
  networking
  registry

  echo "Cluster provisioned and ready to configure"
}

cluster() {
  echo "Creating cluster"
  kind create cluster --config=kind.yaml --wait=60s
}

serving() {
  kubectl apply --filename https://github.com/knative/serving/releases/download/$serving_version/serving-crds.yaml
  sleep 2
  curl -L -s https://github.com/knative/serving/releases/download/$serving_version/serving-core.yaml | yq 'del(.spec.template.spec.containers[]?.resources)' -y | yq 'del(.metadata.annotations."knative.dev/example-checksum")' -y | kubectl apply -f -
  sleep 15
  kubectl get pod -n knative-serving
}

eventing() {
  kubectl apply --filename https://github.com/knative/eventing/releases/download/$eventing_version/eventing-crds.yaml
  sleep 2
  curl -L -s https://github.com/knative/eventing/releases/download/$eventing_version/eventing-core.yaml | yq 'del(.spec.template.spec.containers[]?.resources)' -y | yq 'del(.metadata.annotations."knative.dev/example-checksum")' -y | kubectl apply -f -
  curl -L -s https://github.com/knative/eventing/releases/download/$eventing_version/in-memory-channel.yaml | yq 'del(.spec.template.spec.containers[]?.resources)' -y | yq 'del(.metadata.annotations."knative.dev/example-checksum")' -y | kubectl apply -f -
  curl -L -s https://github.com/knative/eventing/releases/download/$eventing_version/mt-channel-broker.yaml | yq 'del(.spec.template.spec.containers[]?.resources)' -y | yq 'del(.metadata.annotations."knative.dev/example-checksum")' -y | kubectl apply -f -
  sleep 15
  kubectl get pod -n knative-eventing
}

networking() {
  kubectl apply --filename https://github.com/knative-sandbox/net-kourier/releases/download/$kourier_version/kourier.yaml
  kubectl patch configmap/config-network \
      --namespace knative-serving \
      --type merge \
      --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
  sleep 15
  kubectl get pod -n kourier-system
}

registry() {
  kubectl apply -f dev-provision-registry.yaml
  # TODO: - run registry:2 docker image
  #       - connect networks
  #       - flag as insecure (is this possible programatically on darwin?)
  #       - add teardown steps to cleanup script

}

main "$@"
