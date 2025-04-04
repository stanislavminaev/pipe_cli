package jobs

import (
	"fmt"
	"log/slog"
	"os"
	"pipe-cli/internal/config"
	"pipe-cli/internal/jira_wick"

	"github.com/Masterminds/semver/v3"
)

// CloseJiraReleaseVersionJob — структура, реализующая джоб закрытия закрытие релиза в жире
type CloseJiraReleaseVersionJob struct {
	ciVars     *config.GitlabCIVariables
	jobConfig  *config.JobConfig
	jiraClient jira_wick.JiraWickClient
}

// NewCloseJiraReleaseVersionJob — конструктор, возвращающий CloseJiraReleaseVersionJob
func NewCloseJiraReleaseVersionJob() (IPipelineJob, error) {
	return &CloseJiraReleaseVersionJob{
		ciVars:     config.GetGitlabCIVariables(),
		jobConfig:  config.GetJobConfig(),
		jiraClient: jira_wick.NewJiraWickClient(config.GetJobConfig().JiraWickService),
	}, nil
}

// Run - запуск джоба
func (j *CloseJiraReleaseVersionJob) Run() {
	fmt.Println("Running CloseJiraReleaseVersionJob")

	reqBody, err := j.buildRequestPayload()
	if err != nil {
		slog.Error("Failed to build request payload", "error", err)
		os.Exit(1)
	}

	_, err = j.jiraClient.CloseRelease(reqBody)
	if err != nil {
		os.Exit(1)
	}
}

// buildRequestPayload формирует payload для запроса на создание релиза в жире
func (j *CloseJiraReleaseVersionJob) buildRequestPayload() (*jira_wick.CloseReleaseRequest, error) {
	isRelease, err := j.isReleaseVersion()
	if err != nil {
		return nil, err
	}

	deployStatus, err := j.getDeployStatus()
	if err != nil {
		return nil, err
	}

	return &jira_wick.CloseReleaseRequest{
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
