package clients

import (
	"os"
	"os/exec"
	"rpa-git/helpers"
)

type GitClient struct {
	RepositoryPath string `json:"basePath"`
}

func GitCreateClient(path string) *GitClient {
	return &GitClient{
		RepositoryPath: path,
	}
}

func (c GitClient) ExecuteCommand(args ...string) {
	helpers.CheckIfEmpty(c.RepositoryPath, "Invalid repository")

	params := append([]string{"-C", c.RepositoryPath}, args...)

	cmd := exec.Command("git", params...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	helpers.CheckIfError(err)
}
