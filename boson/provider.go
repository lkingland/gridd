package boson

import (
	"fmt"

	boson "github.com/boson-project/func"
	"github.com/boson-project/func/buildpacks"
	"github.com/boson-project/func/docker"
	"github.com/boson-project/func/knative"

	"github.com/lkingland/gridd"
)

const DefaultRegistry = "localhost:5000"

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
	// Should be named Update rather than Deploy
	// Should take a Function rather than root path
	return p.impl.Deploy(f.Root)
}

func (p *Provider) Delete(f gridd.Function) error {
	// Should be named Delete
	return p.impl.Remove(newBosonFunction(f))
}

func (p *Provider) List(f gridd.Function) (names []string, err error) {
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

	deployer, err := knative.NewDeployer("")
	if err != nil {
		panic(err) // TODO: remove error from deployer constructor entirely
	}
	deployer.Verbose = verbose

	return boson.New(
		boson.WithRegistry(registry),
		boson.WithVerbose(verbose),
		boson.WithBuilder(builder),
		boson.WithPusher(pusher),
		boson.WithDeployer(deployer),
	)
}

func newBosonFunction(f gridd.Function) boson.Function {
	return boson.Function{
		Root:    f.Root,
		Name:    f.Name,
		Runtime: f.Language,
	}
}
