package clients

import (
	"os"
	"os/exec"
	"rpa-git/helpers"
	"rpa-git/models"
)

type GitClient struct {
	RepositoriesPath string `json:"basePath"`
}

func GitCreateClient() *GitClient {
	return &GitClient{
		RepositoriesPath: models.RepositoresPath(),
	}
}

func (c GitClient) ExecuteCommand(args ...string) {
	helpers.CheckIfEmpty(c.RepositoriesPath, "Invalid repository")

	params := append([]string{"-C", c.RepositoriesPath}, args...)

	cmd := exec.Command("git", params...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	helpers.CheckIfError(err)
}
