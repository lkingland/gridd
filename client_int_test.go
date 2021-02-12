// +build integration

package gridd_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/lkingland/gridd"
	"github.com/lkingland/gridd/boson"
)

/*
 NOTE:  Running integration tests locally requires a configured test cluster.
        Test failures may require manual removal of dangling resources.

 ## Integration Cluster
 These integration tests require a properly configured cluster,
 such as that which is setup and configured in CI (see .github/workflows).
 A local KinD cluster can be started via:
   ./hack/allocate.sh && ./hack/configure.sh

 ## Integration Testing
 These tests can be run via the make target:
   make integration
	or manually by specifying the tag
   go test -v -tags integration ./...

 ## Teardown and Cleanup
 Tests should clean up after themselves.  In the event of failures, one may
 need to manually remove files:
   rm -rf ./testdata/example.com
 The test cluster is not automatically removed, as it can be reused.  To remove:
   ./hack/delete.sh
*/

func TestList(t *testing.T) {
	verbose := true

	// Assemble
	grid := gridd.New(
		boson.NewProvider(verbose),
		gridd.WithVerbose(verbose))

	// Act
	names, err := grid.List()
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if len(names) != 0 {
		t.Fatalf("Expected no Functions, got %v", names)
	}
}

func TestCreate(t *testing.T) {
	defer within(t, "testdata/example.com/create")()
	verbose := true

	// Assemble
	grid := gridd.New(
		boson.NewProvider(verbose),
		gridd.WithVerbose(verbose))

	// Act
	if err := grid.Create(gridd.Function{}); err != nil {
		t.Fatal(err)
	}
	defer del(t, grid, "create")

	// Assert
	names, err := grid.List()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(names, []string{"create"}) {
		t.Fatalf("Expected function list ['create'], got %v", names)
	}
}

func TestRead(t *testing.T) {
	// TODO
	return
	defer within(t, "testdata/example.com/read")()

	touch("RUNSTAMP-TestRead")
}

func TestUpdate(t *testing.T) {
	defer within(t, "testdata/example.com/update")()
	verbose := true

	grid := gridd.New(
		boson.NewProvider(verbose),
		gridd.WithVerbose(verbose))

	if err := grid.Create(gridd.Function{}); err != nil {
		t.Fatal(err)
	}
	defer del(t, grid, "update")

	if err := grid.Update(gridd.Function{Root: "."}); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	defer within(t, "testdata/example.com/delete")()
	verbose := true

	grid := gridd.New(
		boson.NewProvider(verbose),
		gridd.WithVerbose(verbose))

	if err := grid.Create(gridd.Function{}); err != nil {
		t.Fatal(err)
	}
	waitFor(t, grid, "delete")

	if err := grid.Delete("delete"); err != nil {
		t.Fatal(err)
	}

	names, err := grid.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 0 {
		t.Fatalf("Expected empty Functions list, got %v", names)
	}
}

// Helpers

// Del cleans up after a test by removing a function by name.
// (test fails if the named function does not exist)
//
// Intended to be run in a defer statement immediately after create, del
// works around the asynchronicity of the underlying platform's creation
// step by polling the provider until the names function becomes available
// (or the test times out), before firing off a deletion request.
// Of course, ideally this would be replaced by the use of a synchronous Create
// method, or at a minimum a way to register a callback/listener for the
// creation event.  This is what we have for now, and the show must go on.
func del(t *testing.T, c *gridd.Client, name string) {
	waitFor(t, c, name)
	if err := c.Delete(name); err != nil {
		t.Fatal(err)
	}
}

// waitFor the named Function to become available in List output.
// TODO: the API should be synchronous, but that depends first on
// Create returning the derived name such that we can bake polling in.
// Ideally the Boson provider's Creaet would be made syncrhonous.
func waitFor(t *testing.T, c *gridd.Client, name string) {
	var pollInterval = 2 * time.Second

	for { // ever (i.e. defer to global test timeout)
		nn, err := c.List()
		if err != nil {
			t.Fatal(err)
		}
		for _, n := range nn {
			if n == name {
				return
			}
		}
		time.Sleep(pollInterval)
	}
}

// Create the given directory, CD to it, and return a function which can be
// run in a defer statement to return to the original directory and cleanup.
// Note must be executed, not deferred itself
// NO:  defer within(t, "somedir")
// YES: defer within(t, "somedir")()
func within(t *testing.T, root string) func() {
	cwd := pwd(t)
	mkdir(t, root)
	cd(t, root)
	return func() {
		cd(t, cwd)
		rm(t, root)
	}
}

func pwd(t *testing.T) string {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func mkdir(t *testing.T, dir string) {
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
}

func cd(t *testing.T, dir string) {
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
}

func rm(t *testing.T, dir string) {
	if err := os.RemoveAll(dir); err != nil {
		t.Fatal(err)
	}
}

func touch(file string) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
	t := time.Now().Local()
	if err := os.Chtimes(file, t, t); err != nil {
		panic(err)
	}
}
