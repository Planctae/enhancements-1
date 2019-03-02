package directory

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/planctae/enhancements-tracking-ng/pkg/render"
)

func New(p string, releaseName string) error {
	releaseDir := filepath.Join(p, releaseName)

	err := os.MkdirAll(filepath.Join(releaseDir), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(releaseDir, proposedDirname), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(releaseDir, acceptedDirname), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(releaseDir, rejectedDirname), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(releaseDir, shippedDirname), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(releaseDir, slippedDirname), os.ModePerm)
	if err != nil {
		return err
	}

	ownersBytes, err := render.ReleaseOwners(releaseName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(releaseDir, ownersFilename), ownersBytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

const (
	proposedDirname = "proposed"
	acceptedDirname = "accepted"
	rejectedDirname = "rejected"
	shippedDirname  = "shipped"
	slippedDirname  = "slipped"

	ownersFilename = "OWNERS"
)
