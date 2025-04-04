package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"io"
	"log/slog"
	"net/http"
	"os"
	"pipe-cli/config"
)

// CloseJiraReleaseVersionJob — структура, реализующая джоб закрытия закрытие релиза в жире
type CloseJiraReleaseVersionJob struct {
	ciVariables *config.GitlabCIVariables
	jobConfig   *config.JobConfig
}

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

	response, err := j.sendRequest(reqBody)
	if err != nil {
		slog.Error("Failed to send request", "error", err)
		os.Exit(1)
	}

	if !response.Status {
		slog.Error("Jira responded with error", "comment", response.Comment)
		os.Exit(1)
	}

	slog.Info("Jira response received", "comment", response.Comment)
}

func (j *CloseJiraReleaseVersionJob) buildRequestPayload() ([]byte, error) {
	data := jiraCloseReleaseRequest{
		JiraProjectId:        j.jobConfig.JiraProjectID,
		ReleaseNumber:        j.ciVariables.CommitRefName,
		JiraIssueComponentID: j.jobConfig.JiraIssueComponentID,
		ProjectPath:          j.ciVariables.ProjectPath,
		DeployStatus:         j.getDeployStatus(),
		NotifyChannel:        j.jobConfig.JiraNotifyChannel,
		TechOwners:           j.jobConfig.TechOwners,
		IsRelease:            j.isReleaseVersion(),
	}
	return json.Marshal(data)
}

func (j *CloseJiraReleaseVersionJob) isReleaseVersion() bool {
	if !j.jobConfig.TrunkMode {
		return true
	}

	v, err := semver.NewVersion(j.ciVariables.CommitRefName)
	if err != nil {
		slog.Error("Invalid version format", "version", j.ciVariables.CommitRefName, "error", err)
		os.Exit(1)
	}
	return v.Patch() == 0
}

func (j *CloseJiraReleaseVersionJob) getDeployStatus() bool {
	return true
}

func (j *CloseJiraReleaseVersionJob) sendRequest(data []byte) (*jiraCloseReleaseResponse, error) {
	resp, err := http.Post(j.jobConfig.JiraWickService, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("FLY_ERROR-01-14.04.01: Не удалось подключиться к сервису JiraWick. %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var response jiraCloseReleaseResponse
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Info("Raw response body", "body", string(body))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// NewCloseJiraReleaseVersionJob — конструктор, возвращающий CloseJiraReleaseVersionJob
func NewCloseJiraReleaseVersionJob() (IPipelineJob, error) {
	return &CloseJiraReleaseVersionJob{
		ciVariables: config.GetGitlabCIVariables(),
		jobConfig:   config.GetJobConfig(),
	}, nil
}
