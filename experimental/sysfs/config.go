package sysfs

import (
	"github.com/bananabytelabs/wazero"
	experimentalsys "github.com/bananabytelabs/wazero/experimental/sys"
)

// FSConfig extends wazero.FSConfig, allowing access to the experimental
// sys.FS until it is moved to the "sys" package.
type FSConfig interface {
	// WithSysFSMount assigns a sys.FS file system for any paths beginning at
	// `guestPath`.
	//
	// This is an alternative to WithFSMount, allowing more features.
	WithSysFSMount(fs experimentalsys.FS, guestPath string) wazero.FSConfig
}
