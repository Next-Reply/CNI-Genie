package interfaces

import (
	"io/ioutil"
	"os"
)

type RW interface {
	ReadFile(file string) ([]byte, error)
	ReadDir(dir string) ([]os.FileInfo, error)
	CreateFile(ilePath string, bytes []byte, perm os.FileMode) error
}

type IO struct{}

func (i IO) ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func (i IO) ReadDir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}

func (i IO) CreateFile(filePath string, bytes []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filePath, bytes, perm)
}
