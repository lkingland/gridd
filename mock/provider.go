package mock

import "github.com/lkingland/gridd"

type Provider struct {
	CreateInvoked bool
	CreateFn      func(gridd.Function) error
	ReadInvoked   bool
	ReadFn        func(gridd.Function) (string, error)
	UpdateInvoked bool
	UpdateFn      func(gridd.Function) error
	DeleteInvoked bool
	DeleteFn      func(gridd.Function) error
	ListInvoked   bool
	ListFn        func(gridd.Function) ([]string, error)
}

func NewProvider() *Provider {
	return &Provider{
		CreateFn: func(gridd.Function) error { return nil },
		ReadFn:   func(gridd.Function) (string, error) { return "", nil },
		UpdateFn: func(gridd.Function) error { return nil },
		DeleteFn: func(gridd.Function) error { return nil },
		ListFn:   func(gridd.Function) ([]string, error) { return []string{}, nil },
	}
}

func (p *Provider) Create(f gridd.Function) error {
	p.CreateInvoked = true
	return p.CreateFn(f)
}

func (p *Provider) Read(f gridd.Function) (string, error) {
	p.ReadInvoked = true
	return p.ReadFn(f)
}

func (p *Provider) Update(f gridd.Function) error {
	p.UpdateInvoked = true
	return p.UpdateFn(f)
}

func (p *Provider) Delete(f gridd.Function) error {
	p.DeleteInvoked = true
	return p.DeleteFn(f)
}

func (p *Provider) List(f gridd.Function) ([]string, error) {
	p.ListInvoked = true
	return p.ListFn(f)
}
