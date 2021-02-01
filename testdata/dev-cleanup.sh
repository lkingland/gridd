# Create a local kind cluster with
# Knative Serving, and Kourier networking installed.
# Suitable for use locally during development.
# CI/CD uses the very similar knative-kind action

main() {
  kind delete cluster --name "kind"
  # TODO: remove and recreate docker registry kind-registry
  # see kind-action/cleanup.sh
}

main
