package interfaces

import "os"

type fileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}

type osFileSystem struct{}

func (fs osFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs osFileSystem) Remove(name string) error {
	return os.Remove(name)
}
