package tasks

import (
	"fmt"
	"rpa-git/clients"
	"rpa-git/helpers"
	"rpa-git/models"
	"sync"
)

func GitInitRepos() {
	config := models.AppConfig{}
	config.Load()

	var wg sync.WaitGroup

	for _, repo := range config.Repositories {
		origin_url, name, username := repo.OriginUrl, repo.Name, repo.Auth.Username
		password := helpers.IfThen(repo.Auth.Token, repo.Auth.Password)

		gitea_url, owner, repository := repo.GiteaUrl, repo.GiteaRepo.Owner, repo.GiteaRepo.Repository
		gitea_repo := fmt.Sprintf("%s/%s/%s.git", gitea_url, owner, repository)

		wg.Add(1)
		go func(origin, gitea, name, username, password string) {
			defer wg.Done()

			git := clients.GitCreateClient()
			helpers.Warning("git clone %s", name)

			git.Clone(origin, gitea, name, username, password)
			helpers.Info("%s was cloned", name)
		}(origin_url, gitea_repo, name, username, password)
	}

	wg.Wait()
}
