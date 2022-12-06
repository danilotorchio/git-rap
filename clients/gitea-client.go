package clients

import (
	"rpa-git/helpers"
	"rpa-git/models"

	"code.gitea.io/sdk/gitea"
)

type GiteaClient struct {
	client *gitea.Client
}

func GiteaCreateClient(url string, repo models.GiteaRepository) *GiteaClient {
	client, err := gitea.NewClient(url)
	helpers.CheckIfError(err)

	client.SetBasicAuth(repo.Auth.Username, repo.Auth.Password)

	return &GiteaClient{
		client: client,
	}
}

func (m GiteaClient) CreateBranchCommand(owner, repo, base_branch, new_branch string) {
	_, resp, err := m.client.CreateBranch(owner, repo, gitea.CreateBranchOption{
		OldBranchName: base_branch,
		BranchName:    new_branch,
	})
	defer resp.Body.Close()

	helpers.CheckIfError(err)
}

func (m GiteaClient) ProtectBranchCommand(owner, repo, branch string) {
	_, resp, err := m.client.CreateBranchProtection(owner, repo, gitea.CreateBranchProtectionOption{
		BranchName:                    branch,
		EnablePush:                    true,
		EnablePushWhitelist:           true,
		EnableMergeWhitelist:          true,
		EnableApprovalsWhitelist:      true,
		RequiredApprovals:             1,
		BlockOnRejectedReviews:        true,
		BlockOnOfficialReviewRequests: true,
		DismissStaleApprovals:         true,
		BlockOnOutdatedBranch:         true,
		PushWhitelistTeams:            []string{"Owners"},
		MergeWhitelistTeams:           []string{"Owners"},
		ApprovalsWhitelistTeams:       []string{"Owners"},
	})
	defer resp.Body.Close()

	helpers.CheckIfError(err)
}
