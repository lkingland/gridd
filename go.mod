module github.com/lkingland/gridd

go 1.16

require github.com/boson-project/func v0.12.0

replace (
	// Nail down k8 deps to align with transisitive deps
	k8s.io/client-go => k8s.io/client-go v0.18.12
	k8s.io/code-generator => k8s.io/code-generator v0.18.12
)
