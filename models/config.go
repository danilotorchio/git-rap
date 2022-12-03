package models

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"rpa-git/helpers"
	"runtime"
)

type AuthCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GiteaRepository struct {
	Url        string    `json:"url"`
	Owner      string    `json:"owner"`
	Repository string    `json:"repository"`
	Auth       AuthCreds `json:"auth"`
}

type Repository struct {
	Name      string          `json:"name"`
	RemoteUrl string          `json:"remoteUrl"`
	LocalPath string          `json:"localPath"`
	Auth      AuthCreds       `json:"auth"`
	GiteaRepo GiteaRepository `json:"giteaRepo"`
}

type AppConfig struct {
	Repositories []Repository `json:"repositories"`
}

func AppDir() string {
	user, err := user.Current()
	helpers.CheckIfError(err)

	return fmt.Sprintf("%s/.rpa-git", user.HomeDir)
}

func ConfigFilePath() string {
	return fmt.Sprintf("%s/config.json", AppDir())
}

func InitializeConfig() {
	var config *AppConfig

	if helpers.CheckIfDirectoryExists(AppDir()) {
		config = &AppConfig{}
		config.Load()
	} else {
		config = &AppConfig{
			Repositories: []Repository{
				{
					Name:      "example",
					RemoteUrl: "https://example.com/repo.git",
					LocalPath: "/path/to/repository",
					Auth: AuthCreds{
						Username: "username",
						Password: "password",
					},
					GiteaRepo: GiteaRepository{
						Url:        "http://gitea.example.com",
						Owner:      "organization_name",
						Repository: "repository_name",
						Auth: AuthCreds{
							Username: "username",
							Password: "password",
						},
					},
				},
			},
		}

		config.Save()
	}

	var cmd *exec.Cmd
	params := []string{ConfigFilePath()}

	if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", params...)
	} else {
		cmd = exec.Command("Notepad", params...)
	}

	err := cmd.Run()
	helpers.CheckIfError(err)
}

func (m *AppConfig) Load() {
	file, err := os.ReadFile(ConfigFilePath())
	helpers.CheckIfError(err)

	err = json.Unmarshal(file, &m)
	helpers.CheckIfError(err)
}

func (m *AppConfig) Save() {
	appDir := AppDir()
	helpers.CreateDirectory(appDir)

	filePath := ConfigFilePath()
	fileContent, err := json.MarshalIndent(m, "", "    ")
	helpers.CheckIfError(err)

	f, err := os.Create(filePath)
	helpers.CheckIfError(err)
	defer f.Close()

	_, err = f.Write(append(fileContent, "\n"...))
	helpers.CheckIfError(err)
}
