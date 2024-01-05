package gojs_test

import (
	"testing"

	"github.com/bananabytelabs/wazero"
	"github.com/bananabytelabs/wazero/internal/gojs/config"
	"github.com/bananabytelabs/wazero/internal/testing/require"
)

func Test_argsAndEnv(t *testing.T) {
	t.Parallel()

	stdout, stderr, err := compileAndRun(testCtx, "argsenv", func(moduleConfig wazero.ModuleConfig) (wazero.ModuleConfig, *config.Config) {
		return moduleConfig.WithEnv("c", "d").WithEnv("a", "b"), config.NewConfig()
	})

	require.Zero(t, stderr)
	require.NoError(t, err)
	require.Equal(t, `
args 0 = test
args 1 = argsenv
environ 0 = c=d
environ 1 = a=b
`, stdout)
}
