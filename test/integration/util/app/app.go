package app

import (
	"errors"

	"github.com/onsi/gomega/gexec"
)

func BuildApplication(path string) (string, error) {
	execPath, err := gexec.Build(path)
	if err != nil {
		return "", errors.New("failed to build")
	}
	return execPath, nil
}
