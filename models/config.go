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

type Project struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AppConfig struct {
	Projects []Project `json:"projects"`
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

	if !helpers.CheckIfDirectoryIsEmpty(AppDir()) {
		config = &AppConfig{}
		config.Load()
	} else {
		config = &AppConfig{
			Projects: []Project{
				{
					Name:     "example",
					Url:      "https://example.com/repo.git",
					Username: "cuca",
					Password: "VaiTePegar",
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
