// Licensed under the Apache License, Version 2.0.  See LICENSE file.

// +build !integration

package gridd_test

import (
	"context"
	"testing"

	"github.com/lkingland/gridd"
	"github.com/lkingland/gridd/mock"
)

// TestList ensures the list base case: no errors and an empty list, with the
// underlying provider's List method invoked.
func TestList(t *testing.T) {
	var (
		p = mock.NewProvider()
		c = gridd.New(p)
	)
	names, err := c.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 0 {
		t.Fatalf("unexpected list item received.  Expected 0, got %v", names)
	}
	if !p.ListInvoked {
		t.Fatal("Provider not invoked")
	}
}

// TestCreate ensures the Create base case: no errors and a call to Create
// invokes the underlying provider's Create with no errors.
func TestCreate(t *testing.T) {
	var (
		p = mock.NewProvider()
		c = gridd.New(p)
		f = gridd.Function{}
	)
	if err := c.Create(context.Background(), f); err != nil {
		t.Fatal(err)
	}
	if !p.CreateInvoked {
		t.Fatal("Provider not invoked")
	}
}

// TestCreateDefaultsLanguage ensures that the default language of Gridd is
// provided to the underlying Provider, intended to override its default.
func TestCreateDefaultsLanguage(t *testing.T) {
	var (
		p = mock.NewProvider()
		c = gridd.New(p)
		f = gridd.Function{}
	)
	p.CreateFn = func(ctx context.Context, x gridd.Function) error {
		if x.Language != gridd.DefaultLanguage {
			t.Fatalf("Default language not applied.  Expected '%v' got '%v'",
				gridd.DefaultLanguage, x.Language)
		}
		return nil
	}
	if err := c.Create(context.Background(), f); err != nil {
		t.Fatal(f)
	}
	if !p.CreateInvoked {
		t.Fatal("Provider not invoked")
	}
}

// TestDelete ensures the Update base case: no errors and a call to Delete
// invokes the underlying provider's implementation with no errors.
func TestUpdate(t *testing.T) {
	var (
		p = mock.NewProvider()
		c = gridd.New(p)
	)
	if err := c.Delete(context.Background(), "myfunc"); err != nil {
		t.Fatal(err)
	}
	if !p.DeleteInvoked {
		t.Fatal("Provider not invoked")
	}
}

// TODO TestUpdateEmpty
