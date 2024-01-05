package amd64

import (
	"testing"

	"github.com/bananabytelabs/wazero/internal/asm"
	"github.com/bananabytelabs/wazero/internal/testing/require"
)

// TestInstructionName ensures that all the instruction's name is defined.
func TestInstructionName(t *testing.T) {
	for inst := asm.Instruction(0); inst < instructionEnd; inst++ {
		require.NotEqual(t, "", InstructionName(inst))
	}
}
