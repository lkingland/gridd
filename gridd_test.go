// +build !integration

package gridd_test

import (
	"testing"

	"github.com/lkingland/gridd"
)

// TestUnit is run as a unit test
func TestUnit(t *testing.T) {
	_ = gridd.New()
}
