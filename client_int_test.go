// +build integration

package gridd_test

import (
	"os"
	"testing"

	"github.com/lkingland/gridd"
	"github.com/lkingland/gridd/boson"
)

//
// These integration tests require a properly configured kubernetes cluster,
// such as that which is setup and configured in CI.
//

// TestCreate ensures that creation of a Function with all defaults in an
// empty directory succeeds.
func TestCreate(t *testing.T) {
	// mkdir ...
	root := "testdata/example.com/www" // Root from which to run the test
	if err := os.MkdirAll(root, 0700); err != nil {
		t.Fatal(err)
	}
	// defer os.RemoveAll(root)

	// cd ...
	err := os.Chdir(root)
	if err != nil {
		t.Fatal(err)
	}

	// grid
	verbose := true
	provider := boson.NewProvider(verbose)
	grid := gridd.New(provider, gridd.WithVerbose(verbose))

	// create
	err = grid.Create(gridd.Function{})
	if err != nil {
		t.Fatal(err)
	}

	// TODO: confirm running
}
