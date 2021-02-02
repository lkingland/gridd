package gridd

import (
	"fmt"
	"os"
)

const (
	DefaultLanguage = "go"
	DefaultName     = "www"
)

type Client struct {
	verbose  bool
	registry string
	provider Provider
}

type Function struct {
	Root     string
	Name     string
	Language string
}

type Provider interface {
	Create(Function) error
	Read(Function) (string, error)
	Update(Function) error
	Delete(Function) error
	List(Function) ([]string, error)
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

func (g *Client) Create(f Function) error {
	if f.Root == "" {
		f.Root, _ = os.Getwd()
	}
	if f.Language == "" {
		f.Language = DefaultLanguage
	}
	if f.Name == "" {
		f.Name = DefaultName
	}
	fmt.Printf("Create %#v\n", f)
	return g.provider.Create(f)
}
