package frontend

import (
	"github.com/bananabytelabs/wazero/internal/engine/wazevo/ssa"
	"github.com/bananabytelabs/wazero/internal/wasm"
)

func FunctionIndexToFuncRef(idx wasm.Index) ssa.FuncRef {
	return ssa.FuncRef(idx)
}
