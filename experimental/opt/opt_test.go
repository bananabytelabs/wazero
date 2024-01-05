package opt_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/bananabytelabs/wazero"
	"github.com/bananabytelabs/wazero/experimental/opt"
	"github.com/bananabytelabs/wazero/internal/testing/require"
)

func TestUseOptimizingCompiler(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		return
	}
	c := opt.NewRuntimeConfigOptimizingCompiler()
	r := wazero.NewRuntimeWithConfig(context.Background(), c)
	require.NoError(t, r.Close(context.Background()))
}
