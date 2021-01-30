package gridd_test

import (
	"testing"

	"github.com/lkingland/gridd"
)

// TestNew ensures that the Grid daemon can be instantiated.
func TestNew(t *testing.T) {
	_ = gridd.New()
}
