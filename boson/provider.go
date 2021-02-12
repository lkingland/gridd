package boson

import (
	"fmt"

	boson "github.com/boson-project/func"
	"github.com/boson-project/func/buildpacks"
	"github.com/boson-project/func/docker"
	"github.com/boson-project/func/knative"

	"github.com/lkingland/gridd"
)

const (
	// DefaultRegistry must contain both the registry host and
	// registry namespace at that host until such time as
	// the func project either seprates these values, or provides
	// an in-cluster container registry by default.
	// TODO: open an issue to clarify better; probably best to suggest separating
	// Registry into RegistryHost and RegistryNamespace
	DefaultRegistry = "localhost:5000/gridd"

	// DefaultNamespace for the underlying deployments.  Must be the same
	// as is set up and configured (see hack/configure.sh)
	DefaultNamespace = "func"
)

type Provider struct {
	impl *boson.Client
}

func NewProvider(verbose bool) *Provider {
	return &Provider{impl: newBosonClient(DefaultRegistry, verbose)}
}

func (p *Provider) Create(f gridd.Function) error {
	// This should be Create rather than "New"
	return p.impl.New(newBosonFunction(f))
}

func (p *Provider) Read(f gridd.Function) (string, error) {
	// Should take a func rather thann both name and root
	// ex:  Describe(Function{Name:"A",Root:"B"})
	desc, err := p.impl.Describe(f.Name, f.Root)
	return fmt.Sprintf("%#v", desc), err
}

func (p *Provider) Update(f gridd.Function) error {
	// Should take a Function rather than root path
	// Should probably be named Update rather than Deploy
	return p.impl.Deploy(f.Root)
}

func (p *Provider) Delete(name string) error {
	// Should be named Delete:
	return p.impl.Remove(boson.Function{Name: name})
}

func (p *Provider) List() (names []string, err error) {
	names = []string{}
	// Should be a simple name list, with details retreivable for each lazily.
	ll, err := p.impl.List()
	if err != nil {
		return
	}
	for _, l := range ll {
		names = append(names, l.Name)
	}
	return
}

func newBosonClient(registry string, verbose bool) *boson.Client {
	builder := buildpacks.NewBuilder()
	builder.Verbose = verbose

	pusher := docker.NewPusher()
	pusher.Verbose = verbose

	deployer, err := knative.NewDeployer(DefaultNamespace)
	if err != nil {
		panic(err) // TODO: remove error from deployer constructor
	}
	deployer.Verbose = verbose

	remover, err := knative.NewRemover(DefaultNamespace)
	if err != nil {
		panic(err) // TODO: remove error from remover constructor
	}

	lister, err := knative.NewLister(DefaultNamespace)
	if err != nil {
		panic(err) // TODO: remove error from lister constructor
	}
	lister.Verbose = verbose

	return boson.New(
		boson.WithRegistry(registry),
		boson.WithVerbose(verbose),
		boson.WithBuilder(builder),
		boson.WithPusher(pusher),
		boson.WithDeployer(deployer),
		boson.WithRemover(remover),
		boson.WithLister(lister),
	)
}

func newBosonFunction(f gridd.Function) boson.Function {
	return boson.Function{
		Root:    f.Root,
		Name:    f.Name,
		Runtime: f.Language,
	}
}
