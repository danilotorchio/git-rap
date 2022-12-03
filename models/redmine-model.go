package models

type RedmineIssueProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RedmineIssueTrackerInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RedmineIssueStatusInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsClosed bool   `json:"is_closed"`
}

type RedmineIssuePriorityInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RedmineIssueAuthorInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RedmineIssueCustomFieldInfo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RedmineIssueModel struct {
	Id int `json:"id"`

	Project  RedmineIssueProjectInfo  `json:"project"`
	Tracker  RedmineIssueTrackerInfo  `json:"tracker"`
	Status   RedmineIssueStatusInfo   `json:"status"`
	Priority RedmineIssuePriorityInfo `json:"priority"`
	Author   RedmineIssueAuthorInfo   `json:"author"`

	Subject     string  `json:"subject"`
	Description string  `json:"description"`
	StartDate   string  `json:"start_date"`
	DueDate     string  `json:"due_date"`
	DoneRatio   float32 `json:"done_ratio"`
	IsPrivate   bool    `json:"is_private"`

	Estimated_hours     float32 `json:"estimated_hours"`
	TotalEstimatedHours float32 `json:"total_estimated_hours"`
	SpentHours          float32 `json:"spent_hours"`
	TotalSpentHours     float32 `json:"total_spent_hours"`

	CustomFields []RedmineIssueCustomFieldInfo `json:"custom_fields"`

	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	ClosedOn  string `json:"closed_on"`
}

type RedmineSearchIssuesResponse struct {
	Issues     []RedmineIssueModel `json:"issues"`
	TotalCount int                 `json:"total_count"`
	Offset     int                 `json:"offset"`
	Limit      int                 `json:"limit"`
}
