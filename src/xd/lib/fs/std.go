package fs

import (
	"os"
	"path/filepath"
	"xd/lib/util"
)

type stdFs struct{}

var STD stdFs

func (f stdFs) Open() error {
	return nil
}

func (f stdFs) Close() error {
	return nil
}

func (f stdFs) EnsureDir(fname string) error {
	return util.EnsureDir(fname)
}

func (f stdFs) EnsureFile(fname string, sz uint64) error {
	return util.EnsureFile(fname, sz)
}

func (f stdFs) FileExists(fname string) bool {
	return util.CheckFile(fname)
}

func (f stdFs) Glob(glob string) ([]string, error) {
	return filepath.Glob(glob)
}

func (f stdFs) OpenFileReadOnly(fname string) (ReadFile, error) {
	return os.OpenFile(fname, os.O_RDONLY, 0600)
}

func (f stdFs) OpenFileWriteOnly(fname string) (WriteFile, error) {
	return os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0600)
}

func (f stdFs) RemoveAll(fname string) error {
	return os.RemoveAll(fname)
}

func (f stdFs) Remove(fname string) error {
	return os.Remove(fname)
}

func (f stdFs) Join(parts ...string) string {
	return filepath.Join(parts...)
}

func (f stdFs) Move(oldpath, newpath string) (err error) {
	dir, _ := f.Split(newpath)
	err = f.EnsureDir(dir)
	if err == nil {
		err = os.Rename(oldpath, newpath)
	}
	return
}

func (f stdFs) Split(path string) (base, file string) {
	base, file = filepath.Split(path)
	return
}
