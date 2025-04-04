package jira_wick

// CloseReleaseRequest структура для отправки API запроса к сервису JiraWick
type CloseReleaseRequest struct {
	JiraProjectId        string `json:"jira_project_id"`
	ReleaseNumber        string `json:"release_number"`
	JiraIssueComponentID string `json:"jira_issue_component_id"`
	ProjectPath          string `json:"project_path"`
	DeployStatus         bool   `json:"deploy_status"`
	NotifyChannel        string `json:"notify_channel"`
	TechOwners           string `json:"tech_owners"`
	IsRelease            bool   `json:"is_release"`
}

// CloseReleaseResponse структура ответа на запрос создания релиза от JiraWick
type CloseReleaseResponse struct {
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
}
