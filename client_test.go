// +build !integration

package gridd_test

import (
	"testing"

	"github.com/lkingland/gridd"
	"github.com/lkingland/gridd/mock"
)

func TestCreate(t *testing.T) {
	var (
		p = mock.NewProvider()
		c = gridd.New(p)
		f = gridd.Function{}
	)
	if err := c.Create(f); err != nil {
		t.Fatal(f)
	}
	if !p.CreateInvoked {
		t.Fatal("Provider not invoked")
	}
}
