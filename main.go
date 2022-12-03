package main

import (
	"os"
	"rpa-git/helpers"
	"rpa-git/models"
	"rpa-git/tasks"

	"golang.org/x/exp/slices"
)

func main() {
	helpers.CheckArgs("<op = init | run>")
	op := os.Args[1]

	switch op {
	case "init":
		models.InitializeConfig()
	case "run":
		helpers.CheckArgs("<op>", "<project>", "<branch_origin>", "<task_type>", "<task_ref>")
		project, branch_origin, task_type, task_ref := os.Args[2], os.Args[3], os.Args[4], os.Args[5]

		Execute(project, branch_origin, task_type, task_ref)
	default:
		helpers.Warning("Invalid option")
	}

	os.Exit(0)
}

func Execute(project, branch_origin, task_type, task_ref string) {
	config := models.AppConfig{}
	config.Load()

	idx := slices.IndexFunc(config.Repositories, func(c models.Repository) bool { return c.Name == project })
	if idx < 0 {
		helpers.Warning("Project not founded")
		os.Exit(1)
	}

	pConf := config.Repositories[idx]

	tasks.GiteaCreateDevBranchs(pConf, branch_origin, task_type, task_ref)
	tasks.RedmineIssueUpdateDevBranchs(pConf, branch_origin, task_type, task_ref)
}
