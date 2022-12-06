package tasks

import (
	"rpa-git/clients"
	"rpa-git/helpers"
	"rpa-git/models"
	"strings"
	"sync"
)

func GitInitRepos() {
	config := models.AppConfig{}
	config.Load()

	var wg sync.WaitGroup

	for _, repo := range config.Repositories {
		origin, gitea, name, username := repo.OriginUrl, repo.GiteaUrl, repo.Name, repo.Auth.Username
		var password string

		if strings.TrimSpace(repo.Auth.Token) != "" {
			password = repo.Auth.Token
		} else {
			password = repo.Auth.Password
		}

		wg.Add(1)
		go func(origin, gitea, name, username, password string) {
			defer wg.Done()

			git := clients.GitCreateClient()
			helpers.Warning("git clone %s", name)

			git.ExecuteCloneCommand(origin, gitea, name, username, password)
			helpers.Info("%s was cloned", name)
		}(origin, gitea, name, username, password)
	}

	wg.Wait()
}
