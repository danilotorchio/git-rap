package tasks

import (
	"fmt"
	"rpa-git/clients"
	"rpa-git/helpers"
	"rpa-git/models"
)

func GiteaCreateDevBranchs(conf models.Repository, branch_origin, task_type, task_ref string) {
	project_name := conf.Name

	git := clients.GitCreateClient()
	gitea := clients.GiteaCreateClient(conf.GiteaUrl, conf.GiteaRepo)

	helpers.Info("git checkout %s", branch_origin)
	git.Checkout(project_name, branch_origin, conf.Auth)

	helpers.Info("git pull %s %s", "origin", branch_origin)
	git.Pull(project_name, "origin", conf.Auth)

	helpers.Info("git push %s %s", "gitea", branch_origin)
	git.Push(project_name, "gitea", conf.GiteaRepo.Auth)

	branch_dv := fmt.Sprintf("dev/%s", task_ref)
	branch_pr := fmt.Sprintf("%s/%s", task_type, task_ref)

	owner, repo := conf.GiteaRepo.Owner, conf.GiteaRepo.Repository

	helpers.Info("Create dev branch (%s)", branch_dv)
	gitea.CreateBranchCommand(owner, repo, branch_origin, branch_dv)

	helpers.Info("Create PR branch (%s)", branch_pr)
	gitea.CreateBranchCommand(owner, repo, branch_origin, branch_pr)

	helpers.Info("Protect PR branch (%s)", branch_pr)
	gitea.ProtectBranchCommand(owner, repo, branch_pr)
}
