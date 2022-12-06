package clients

import (
	"fmt"
	"os"
	"os/exec"
	"rpa-git/helpers"
	"rpa-git/models"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type GitClient struct {
	BasePath string
}

func GitCreateClient() *GitClient {
	return &GitClient{
		BasePath: models.RepositoriesDir(),
	}
}

func (c GitClient) ExecuteCommand(args ...string) {
	helpers.CheckIfEmpty(c.BasePath, "Invalid repository")

	params := append([]string{"-C", c.BasePath}, args...)

	cmd := exec.Command("git", params...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	helpers.CheckIfError(err)
}

func (c GitClient) Clone(origin, gitea, project_name, username, password string) {
	directory := fmt.Sprintf("%s/%s", models.RepositoriesDir(), project_name)

	if !helpers.CheckIfDirectoryExists(directory) {
		r, err := git.PlainClone(directory, false, &git.CloneOptions{
			URL: origin,
			Auth: &http.BasicAuth{
				Username: username,
				Password: password,
			},
		})
		helpers.CheckIfError(err)

		_, err = r.CreateRemote(&config.RemoteConfig{
			Name: "gitea",
			URLs: []string{gitea},
		})
		helpers.CheckIfError(err)
	}
}

func (c GitClient) Checkout(project_name, branch string, creds models.AuthCreds) {
	directory := fmt.Sprintf("%s/%s", models.RepositoriesDir(), project_name)

	r, err := git.PlainOpen(directory)
	helpers.CheckIfError(err)

	w, err := r.Worktree()
	helpers.CheckIfError(err)

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Force:  true,
	})
	helpers.CheckIfError(err)
}

func (c GitClient) Pull(project_name, remote string, creds models.AuthCreds) {
	directory := fmt.Sprintf("%s/%s", models.RepositoriesDir(), project_name)

	r, err := git.PlainOpen(directory)
	helpers.CheckIfError(err)

	w, err := r.Worktree()
	helpers.CheckIfError(err)

	err = w.Pull(&git.PullOptions{
		RemoteName: remote,
		Auth: &http.BasicAuth{
			Username: creds.Username,
			Password: helpers.IfThen(creds.Token, creds.Password),
		},
	})

	if err != git.NoErrAlreadyUpToDate {
		helpers.CheckIfError(err)
	}
}

func (c GitClient) Push(project_name, remote string, creds models.AuthCreds) {
	directory := fmt.Sprintf("%s/%s", models.RepositoriesDir(), project_name)

	r, err := git.PlainOpen(directory)
	helpers.CheckIfError(err)

	err = r.Push(&git.PushOptions{
		RemoteName: remote,
		Auth: &http.BasicAuth{
			Username: creds.Username,
			Password: helpers.IfThen(creds.Token, creds.Password),
		},
	})

	if err != git.NoErrAlreadyUpToDate {
		helpers.CheckIfError(err)
	}
}
