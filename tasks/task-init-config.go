package tasks

import (
	"os/exec"
	"rpa-git/helpers"
	"rpa-git/models"
	"runtime"
)

func InitializeConfig() {
	var config *models.AppConfig

	if helpers.CheckIfDirectoryExists(models.AppDir()) {
		config = &models.AppConfig{}
		config.Load()
	} else {
		config = &models.AppConfig{
			Repositories: []models.Repository{
				{
					Name:      "example",
					OriginUrl: "https://domain.com/repo.git",
					GiteaUrl:  "http://gitea.domain.com",
					Auth: models.AuthCreds{
						Username: "username",
						Password: "password",
						Token:    "token",
					},
					GiteaRepo: models.GiteaRepository{
						Owner:      "Organization name",
						Repository: "Repository name",
						Auth: models.AuthCreds{
							Username: "username",
							Password: "password",
							Token:    "token",
						},
					},
				},
			},
		}

		config.Save()
	}

	var cmd *exec.Cmd
	params := []string{models.ConfigFilePath()}

	if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", params...)
	} else {
		cmd = exec.Command("Notepad", params...)
	}

	err := cmd.Run()
	helpers.CheckIfError(err)
}
