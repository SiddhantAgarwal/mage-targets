package gobuild

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

func (Go) Build() error {
	return sh.Run("go", "build", "-v", "-o", "./bin", "./...")
}

func (Go) Test() error {
	return sh.Run("go", "test", "-race", "-v", "./...")
}

func (Go) Fmt() error {
	return sh.Run("go", "fmt", "./...")
}

func (Go) Vet() error {
	return sh.Run("go", "vet", "-v", "./...")
}

func (Go) Tidy() error {
	return sh.Run("go", "mod", "tidy", "-v")
}
