module github.com/lkingland/gridd

go 1.15

require github.com/boson-project/func v0.11.0

replace (
	// Nail down k8 deps to align with transisitive deps
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
