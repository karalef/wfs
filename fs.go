package execfs

import (
	"io/fs"
	"os"
	"path/filepath"
)

// AsFS returns the file system rooted in the exec directory.
func AsFS() fs.FS {
	return GetFS("")
}

// GetFS returns the file system rooted relavite to exec directory.
func GetFS(name string) fs.FS {
	return os.DirFS(filepath.Join(dir, name))
}

// ReadDir reads the related named directory and returns a list of directory entries.
func ReadDir(name string) ([]fs.DirEntry, error) {
	f, err := OpenOS(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.ReadDir(-1)
}