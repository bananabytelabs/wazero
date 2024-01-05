//go:build !windows && !plan9

package sysfs

import (
	"syscall"

	"github.com/bananabytelabs/wazero/experimental/sys"
)

func rename(from, to string) sys.Errno {
	if from == to {
		return 0
	}
	return sys.UnwrapOSError(syscall.Rename(from, to))
}
