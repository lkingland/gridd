// Licensed under the Apache License, Version 2.0.  See LICENSE file.
package gridd

import "context"

const DefaultLanguage = "go"

type Client struct {
	verbose  bool
	provider Provider
}

type Function struct {
	Root     string
	Name     string
	Language string
}

type Provider interface {
	Create(context.Context, Function) error
	Read(context.Context, Function) (string, error)
	Update(context.Context, Function) error
	Delete(context.Context, string) error
	List(context.Context) ([]string, error)
}

type Option func(*Client)

func WithVerbose(v bool) Option {
	return func(g *Client) {
		g.verbose = v
	}
}

func New(provider Provider, options ...Option) *Client {
	g := &Client{
		provider: provider,
	}
	for _, o := range options {
		o(g)
	}
	return g
}

func (g *Client) List(ctx context.Context) ([]string, error) {
	return g.provider.List(ctx)
}

func (g *Client) Create(ctx context.Context, f Function) error {
	// The only default value overridden by this library is
	// to presume Go as the default language rather than
	// the Func project's Node.js.
	if f.Language == "" {
		f.Language = DefaultLanguage
	}
	// TODO: should this be the default of the client library?
	if f.Root == "" {
		f.Root = "."
	}

	// TODO: the provider's Create should be synchronous.
	// Emulate this behavior by polling until it is available or timeout.
	// But this can not be done until such time as the provider's Create
	// method returns a populated Function object with a name for which
	// to check.
	return g.provider.Create(ctx, f)
}

func (g *Client) Update(ctx context.Context, f Function) error {
	return g.provider.Update(ctx, f)
}

func (g *Client) Delete(ctx context.Context, name string) error {
	return g.provider.Delete(ctx, name)
}
