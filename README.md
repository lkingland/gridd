# Grid

Proof-of-concept Boson Func client for use in rapid prototyping.

Run `make help` for available actions.

## Building


## Local Cluster Requirement

The Grid Daemon (gridd) requires an available Kubernetes-compliant cluster with
Knative installed and a few configuration modifications.  For local development
and testing, a KinD (Kubernetes in Docker) cluster can be used by following the
below procedure.

Note that the scripts referenced below assume there is not already a cluster 
named `kind`, network named `kind`, or docker process named `kind-registry`.

1.  Prerequisites
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- kubectl
- yq (note on OSX `brew install python-yq`

2.  Allocate Cluster Resources
`hack/allocate.sh`

Sets up a two-node cluster with Knative, Kourier, and a local container registry.
Note that 'kind-registry' must be added manually to the local /etc/hosts file,
and to the docker daemon.config file as an insecure registry.  See the script's
final output for details.

3.  Configure Cluster
`hack/configure.sh`

Configures the cluster for hosting Functions, including namespace creation and setup.

Use `delete.sh` to teardown.

The cluster is now available for local integration and acceptance testing.





