package util

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func IsDir(dirInput string) bool {
	fi, err := os.Stat(dirInput)
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}

	return false
}

func CreateDir(folderPath string) error {
	if IsDir(folderPath) {
		return nil
	}
	return os.MkdirAll(folderPath, os.ModePerm)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func InitCliHomeDir(templateDir string) error {
	if !IsDir(templateDir) {
		return DefaultLoading(func() error {
			return CreateDir(templateDir)
		},
			fmt.Sprintf("'%s' does not exists or not a directory, creating it now...", templateDir),
			":robot:",
		)
	}
	return nil
}

func CloneAndPullTemplatesRepo(repo string, templateDir string) error {
	return DefaultLoading(func() error {
		// clone
		repository, err := git.PlainClone(templateDir, false, &git.CloneOptions{
			URL: repo,
		})
		if err != nil && err != git.ErrRepositoryAlreadyExists {
			return err
		}

		if repository == nil {
			repository, err = git.PlainOpen(templateDir)
			if err != nil {
				return err
			}
		}

		// pull latest changes
		wt, err := repository.Worktree()
		if err != nil {
			return err
		}

		pullErr := wt.Pull(&git.PullOptions{Force: true})
		if pullErr != nil && pullErr != git.NoErrAlreadyUpToDate {
			return pullErr
		}

		return nil
	},
		fmt.Sprintf("Getting latest changes from templates repository (%s)...", repo),
		":robot:",
	)
}
