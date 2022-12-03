package main

import (
	"fmt"
	"os"
	"rpa-git/helpers"
	"rpa-git/models"
	"rpa-git/tasks"

	"golang.org/x/exp/slices"
)

func main() {
	helpers.CheckArgs("<op = help | config | init | run>")
	op := os.Args[1]

	switch op {
	case "help":
		Help()
	case "config":
		InitializeConfig()
	case "init":
		InitializeRepos()
	case "run":
		helpers.CheckArgs("<op>", "<project>", "<branch_origin>", "<task_type>", "<task_ref>")
		project, branch_origin, task_type, task_ref := os.Args[2], os.Args[3], os.Args[4], os.Args[5]

		Run(project, branch_origin, task_type, task_ref)
	default:
		helpers.Warning("Invalid option")
	}

	os.Exit(0)
}

func Help() {
	helpers.Warning("\nAvaliable options:")

	fmt.Printf("\n- %sDisplay this help section.\n", helpers.GetInfo("help"))
	fmt.Printf("\n- %sInitialize de app config file.\n", helpers.GetInfo("config"))
	fmt.Printf("\n- %sInitialize the repositories configured in the app config file.\n", helpers.GetInfo("init"))
	fmt.Printf("\n- %sRun the automation.\n", helpers.GetInfo("run"))
}

func InitializeConfig() {
	models.InitializeConfig()
}

func InitializeRepos() {
	tasks.GitInitRepos()
}

func Run(project, branch_origin, task_type, task_ref string) {
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
