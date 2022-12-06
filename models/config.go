package models

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"rpa-git/helpers"
)

type AuthCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type GiteaRepository struct {
	Owner      string    `json:"owner"`
	Repository string    `json:"repository"`
	Auth       AuthCreds `json:"auth"`
}

type Repository struct {
	Name      string          `json:"name"`
	OriginUrl string          `json:"originUrl"`
	GiteaUrl  string          `json:"giteaUrl"`
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

func RepositoresPath() string {
	return fmt.Sprintf("%s/repositories", AppDir())
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
