package spectest

import (
	"context"
	"embed"
	"testing"

	"github.com/bananabytelabs/wazero"
	"github.com/bananabytelabs/wazero/api"
	"github.com/bananabytelabs/wazero/internal/integration_test/spectest"
	"github.com/bananabytelabs/wazero/internal/platform"
)

//go:embed testdata/*.wasm
//go:embed testdata/*.json
var testcases embed.FS

const enabledFeatures = api.CoreFeaturesV2 // TODO: Enable threads feature after implementing interpreter support

func TestCompiler(t *testing.T) {
	t.Skip() // TODO: Delete after implementing compiler support
	if !platform.CompilerSupported() {
		t.Skip()
	}
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigCompiler().WithCoreFeatures(enabledFeatures))
}

func TestInterpreter(t *testing.T) {
	t.Skip() // TODO: Delete after implementing interpreter support
	spectest.Run(t, testcases, context.Background(), wazero.NewRuntimeConfigInterpreter().WithCoreFeatures(enabledFeatures))
}
