package tasks

import (
	"fmt"
	"rpa-git/clients"
	"rpa-git/helpers"
	"rpa-git/models"
)

func GiteaCreateDevBranchs(conf models.Repository, branch_origin, task_type, task_ref string) {
	git := clients.GitCreateClient(conf.LocalPath)
	gitea := clients.GiteaCreateClient(conf.GiteaRepo)

	helpers.Info("git checkout %s", branch_origin)
	git.ExecuteCommand("checkout", branch_origin)

	helpers.Info("git pull origin %s", branch_origin)
	git.ExecuteCommand("pull", "origin", branch_origin)

	helpers.Info("git push adapter %s", branch_origin)
	git.ExecuteCommand("push", "adapter", branch_origin)

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
