package clients

import (
	"fmt"
	"os"
	"os/exec"
	"rpa-git/helpers"
)

type GitClient struct {
	BasePath string `json:"basePath"`
}

func (c GitClient) ExecuteCommand(project string, args ...string) {
	repository := fmt.Sprintf("%s/%s", c.BasePath, project)
	helpers.CheckIfEmpty(repository, "Invalid repository")

	params := append([]string{"-C", repository}, args...)

	cmd := exec.Command("git", params...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	helpers.CheckIfError(err)
}

func (c GitClient) Initialize() {
	helpers.CheckIfEmpty(c.BasePath, "Invalid repositories path")

}
