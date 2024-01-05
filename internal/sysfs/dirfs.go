package sysfs

import (
	"io/fs"
	"os"

	experimentalsys "github.com/bananabytelabs/wazero/experimental/sys"
	"github.com/bananabytelabs/wazero/internal/platform"
	"github.com/bananabytelabs/wazero/sys"
)

func DirFS(dir string) experimentalsys.FS {
	return &dirFS{
		dir:        dir,
		cleanedDir: ensureTrailingPathSeparator(dir),
	}
}

func ensureTrailingPathSeparator(dir string) string {
	if !os.IsPathSeparator(dir[len(dir)-1]) {
		return dir + string(os.PathSeparator)
	}
	return dir
}

// dirFS is not exported because the input fields must be maintained together.
// This is likely why os.DirFS doesn't, either!
type dirFS struct {
	experimentalsys.UnimplementedFS

	dir string
	// cleanedDir is for easier OS-specific concatenation, as it always has
	// a trailing path separator.
	cleanedDir string
}

// String implements fmt.Stringer
func (d *dirFS) String() string {
	return d.dir
}

// OpenFile implements the same method as documented on sys.FS
func (d *dirFS) OpenFile(path string, flag experimentalsys.Oflag, perm fs.FileMode) (experimentalsys.File, experimentalsys.Errno) {
	return OpenOSFile(d.join(path), flag, perm)
}

// Lstat implements the same method as documented on sys.FS
func (d *dirFS) Lstat(path string) (sys.Stat_t, experimentalsys.Errno) {
	return lstat(d.join(path))
}

// Stat implements the same method as documented on sys.FS
func (d *dirFS) Stat(path string) (sys.Stat_t, experimentalsys.Errno) {
	return stat(d.join(path))
}

// Mkdir implements the same method as documented on sys.FS
func (d *dirFS) Mkdir(path string, perm fs.FileMode) (errno experimentalsys.Errno) {
	err := os.Mkdir(d.join(path), perm)
	if errno = experimentalsys.UnwrapOSError(err); errno == experimentalsys.ENOTDIR {
		errno = experimentalsys.ENOENT
	}
	return
}

// Chmod implements the same method as documented on sys.FS
func (d *dirFS) Chmod(path string, perm fs.FileMode) experimentalsys.Errno {
	err := os.Chmod(d.join(path), perm)
	return experimentalsys.UnwrapOSError(err)
}

// Rename implements the same method as documented on sys.FS
func (d *dirFS) Rename(from, to string) experimentalsys.Errno {
	from, to = d.join(from), d.join(to)
	return rename(from, to)
}

// Readlink implements the same method as documented on sys.FS
func (d *dirFS) Readlink(path string) (string, experimentalsys.Errno) {
	// Note: do not use syscall.Readlink as that causes race on Windows.
	// In any case, syscall.Readlink does almost the same logic as os.Readlink.
	dst, err := os.Readlink(d.join(path))
	if err != nil {
		return "", experimentalsys.UnwrapOSError(err)
	}
	return platform.ToPosixPath(dst), 0
}

// Link implements the same method as documented on sys.FS
func (d *dirFS) Link(oldName, newName string) experimentalsys.Errno {
	err := os.Link(d.join(oldName), d.join(newName))
	return experimentalsys.UnwrapOSError(err)
}

// Rmdir implements the same method as documented on sys.FS
func (d *dirFS) Rmdir(path string) experimentalsys.Errno {
	return rmdir(d.join(path))
}

// Unlink implements the same method as documented on sys.FS
func (d *dirFS) Unlink(path string) (err experimentalsys.Errno) {
	return unlink(d.join(path))
}

// Symlink implements the same method as documented on sys.FS
func (d *dirFS) Symlink(oldName, link string) experimentalsys.Errno {
	// Note: do not resolve `oldName` relative to this dirFS. The link result is always resolved
	// when dereference the `link` on its usage (e.g. readlink, read, etc).
	// https://github.com/bytecodealliance/cap-std/blob/v1.0.4/cap-std/src/fs/dir.rs#L404-L409
	err := os.Symlink(oldName, d.join(link))
	return experimentalsys.UnwrapOSError(err)
}

// Utimens implements the same method as documented on sys.FS
func (d *dirFS) Utimens(path string, atim, mtim int64) experimentalsys.Errno {
	return utimens(d.join(path), atim, mtim)
}

func (d *dirFS) join(path string) string {
	switch path {
	case "", ".", "/":
		if d.cleanedDir == "/" {
			return "/"
		}
		// cleanedDir includes an unnecessary delimiter for the root path.
		return d.cleanedDir[:len(d.cleanedDir)-1]
	}
	// TODO: Enforce similar to safefilepath.FromFS(path), but be careful as
	// relative path inputs are allowed. e.g. dir or path == ../
	return d.cleanedDir + path
}
