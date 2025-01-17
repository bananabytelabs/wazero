//go:build !linux && !darwin && !windows

package sysfs

import (
	"net"
	"syscall"

	"github.com/bananabytelabs/wazero/experimental/sys"
	experimentalsys "github.com/bananabytelabs/wazero/experimental/sys"
	"github.com/bananabytelabs/wazero/internal/fsapi"
	socketapi "github.com/bananabytelabs/wazero/internal/sock"
)

// MSG_PEEK is a filler value.
const MSG_PEEK = 0x2

func newTCPListenerFile(tl *net.TCPListener) socketapi.TCPSock {
	return &unsupportedSockFile{}
}

type unsupportedSockFile struct {
	baseSockFile
}

// Accept implements the same method as documented on socketapi.TCPSock
func (f *unsupportedSockFile) Accept() (socketapi.TCPConn, sys.Errno) {
	return nil, sys.ENOSYS
}

func _pollSock(conn syscall.Conn, flag fsapi.Pflag, timeoutMillis int32) (bool, sys.Errno) {
	return false, sys.ENOTSUP
}

func setNonblockSocket(fd uintptr, enabled bool) sys.Errno {
	return sys.ENOTSUP
}

func readSocket(fd uintptr, buf []byte) (int, sys.Errno) {
	return -1, sys.ENOTSUP
}

func writeSocket(fd uintptr, buf []byte) (int, sys.Errno) {
	return -1, sys.ENOTSUP
}

func recvfrom(fd uintptr, buf []byte, flags int32) (n int, errno sys.Errno) {
	return -1, sys.ENOTSUP
}

// syscallConnControl extracts a syscall.RawConn from the given syscall.Conn and applies
// the given fn to a file descriptor, returning an integer or a nonzero syscall.Errno on failure.
//
// syscallConnControl streamlines the pattern of extracting the syscall.Rawconn,
// invoking its syscall.RawConn.Control method, then handling properly the errors that may occur
// within fn or returned by syscall.RawConn.Control itself.
func syscallConnControl(conn syscall.Conn, fn func(fd uintptr) (int, experimentalsys.Errno)) (n int, errno sys.Errno) {
	return -1, sys.ENOTSUP
}
