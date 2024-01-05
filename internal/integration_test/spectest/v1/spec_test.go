package v1

import (
	"context"
	"runtime"
	"testing"

	"github.com/bananabytelabs/wazero"
	"github.com/bananabytelabs/wazero/api"
	"github.com/bananabytelabs/wazero/experimental/opt"
	"github.com/bananabytelabs/wazero/internal/integration_test/spectest"
	"github.com/bananabytelabs/wazero/internal/platform"
)

func TestCompiler(t *testing.T) {
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestInterpreter(t *testing.T) {
	spectest.Run(t, Testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(api.CoreFeaturesV1))
}

func TestWazevo(t *testing.T) {
	if runtime.GOARCH != "arm64" {
		t.Skip()
	}
	c := opt.NewRuntimeConfigOptimizingCompiler().WithCoreFeatures(api.CoreFeaturesV1)
	spectest.Run(t, Testcases, context.Background(), c)
}
