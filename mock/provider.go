// Licensed under the Apache License, Version 2.0.  See LICENSE file.
package mock

import (
	"context"
	"github.com/lkingland/gridd"
)

type Provider struct {
	CreateInvoked bool
	CreateFn      func(context.Context, gridd.Function) error
	ReadInvoked   bool
	ReadFn        func(context.Context, gridd.Function) (string, error)
	UpdateInvoked bool
	UpdateFn      func(context.Context, gridd.Function) error
	DeleteInvoked bool
	DeleteFn      func(context.Context, string) error
	ListInvoked   bool
	ListFn        func(context.Context) ([]string, error)
}

func NewProvider() *Provider {
	return &Provider{
		CreateFn: func(context.Context, gridd.Function) error { return nil },
		ReadFn:   func(context.Context, gridd.Function) (string, error) { return "", nil },
		UpdateFn: func(context.Context, gridd.Function) error { return nil },
		DeleteFn: func(context.Context, string) error { return nil },
		ListFn:   func(context.Context) ([]string, error) { return []string{}, nil },
	}
}

func (p *Provider) Create(ctx context.Context, f gridd.Function) error {
	p.CreateInvoked = true
	return p.CreateFn(ctx, f)
}

func (p *Provider) Read(ctx context.Context, f gridd.Function) (string, error) {
	p.ReadInvoked = true
	return p.ReadFn(ctx, f)
}

func (p *Provider) Update(ctx context.Context, f gridd.Function) error {
	p.UpdateInvoked = true
	return p.UpdateFn(ctx, f)
}

func (p *Provider) Delete(ctx context.Context, name string) error {
	p.DeleteInvoked = true
	return p.DeleteFn(ctx, name)
}

func (p *Provider) List(ctx context.Context) ([]string, error) {
	p.ListInvoked = true
	return p.ListFn(ctx)
}
