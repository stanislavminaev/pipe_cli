package jobs

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"os"
	"pipe-cli/internal/config"
)

// CloseJiraReleaseVersionJob — структура, реализующая джоб закрытия закрытие релиза в жире
type CloseJiraReleaseVersionJob struct {
	ciVars    *config.GitlabCIVariables
	jobConfig *config.JobConfig
}

// jiraCloseReleaseRequest структура для отправки API запроса к сервису JiraWick
type jiraCloseReleaseRequest struct {
	JiraProjectId        string `json:"jira_project_id"`
	ReleaseNumber        string `json:"release_number"`
	JiraIssueComponentID string `json:"jira_issue_component_id"`
	ProjectPath          string `json:"project_path"`
	DeployStatus         bool   `json:"deploy_status"`
	NotifyChannel        string `json:"notify_channel"`
	TechOwners           string `json:"tech_owners"`
	IsRelease            bool   `json:"is_release"`
}

// jiraCloseReleaseResponse структура ответа на запрос создания релиза от JiraWick
type jiraCloseReleaseResponse struct {
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
}

// Run - запуск джоба
func (j *CloseJiraReleaseVersionJob) Run() {
	fmt.Println("Running CloseJiraReleaseVersionJob")

	reqBody, err := j.buildRequestPayload()
	if err != nil {
		slog.Error("Failed to build request payload", "error", err)
		os.Exit(1)
	}

	response := j.sendRequest(reqBody)

	slog.Info("Jira response received", "comment", response.Comment)
}

// buildRequestPayload формирует payload для запроса на создание релиза в жире
func (j *CloseJiraReleaseVersionJob) buildRequestPayload() (*jiraCloseReleaseRequest, error) {
	isRelease, err := j.isReleaseVersion()
	if err != nil {
		return nil, err
	}

	deployStatus, err := j.getDeployStatus()
	if err != nil {
		return nil, err
	}

	return &jiraCloseReleaseRequest{
		JiraProjectId:        j.jobConfig.JiraProjectID,
		ReleaseNumber:        j.ciVars.CommitRefName,
		JiraIssueComponentID: j.jobConfig.JiraIssueComponentID,
		ProjectPath:          j.ciVars.ProjectPath,
		DeployStatus:         deployStatus,
		NotifyChannel:        j.jobConfig.JiraNotifyChannel,
		TechOwners:           j.jobConfig.TechOwners,
		IsRelease:            isRelease,
	}, nil
}

// isReleaseVersion признак нового релиза
func (j *CloseJiraReleaseVersionJob) isReleaseVersion() (bool, error) {
	if !j.jobConfig.TrunkMode {
		return true, nil
	}

	v, err := semver.NewVersion(j.ciVars.CommitRefName)
	if err != nil {
		return false, fmt.Errorf("invalid version format: %s error: %w", j.ciVars.CommitRefName, err)
	}
	return v.Patch() == 0, nil
}

// getDeployStatus статус деплоя на прод
func (j *CloseJiraReleaseVersionJob) getDeployStatus() (bool, error) {
	return true, nil
}

// sendRequest отправка запроса на создание релиза в жире
func (j *CloseJiraReleaseVersionJob) sendRequest(request *jiraCloseReleaseRequest) *jiraCloseReleaseResponse {

	client := resty.New()
	var response jiraCloseReleaseResponse
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(j.jobConfig.JiraWickService)

	if err != nil {
		slog.Error("Не удалось подключиться к сервису JiraWick", "error", err)
		os.Exit(1)
	}

	if !response.Status {
		slog.Error("Ошибка JiraWick", "comment", response.Comment)
		os.Exit(1)
	}

	slog.Info("Ответ JiraWick", "comment", response.Comment)
	return &response
}

// NewCloseJiraReleaseVersionJob — конструктор, возвращающий CloseJiraReleaseVersionJob
func NewCloseJiraReleaseVersionJob() (IPipelineJob, error) {
	return &CloseJiraReleaseVersionJob{
		ciVars:    config.GetGitlabCIVariables(),
		jobConfig: config.GetJobConfig(),
	}, nil
}
