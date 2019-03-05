package git

import (
	"io"

	libgit "gopkg.in/src-d/go-git.v4"
)

func Clone(repoUrl string, toLocation string, progress io.Writer) error {
	_, err := libgit.PlainClone(toLocation, false, &libgit.CloneOptions{
		URL:      repoUrl,
		Progress: progress,
	})

	return err
}
