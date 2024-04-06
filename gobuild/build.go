package gobuild

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

func (Go) Build() error {
	return sh.Run("go", "build", "-o", "./bin", "./...")
}

func (Go) Test() error {
	return sh.Run("go", "test", "-race", "-v", "./...")
}

func (Go) Fmt() error {
	return sh.Run("go", "fmt", "./...")
}
