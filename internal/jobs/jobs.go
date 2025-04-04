package jobs

import (
	"log/slog"
	"os"
)

// IPipelineJob — интерфейс, который должны реализовать все джобы
type IPipelineJob interface {
	Run()
}

// JobFactory — функция-конструктор для PipelineJob
type JobFactory func() (IPipelineJob, error)

// GetJob - возвращает джоб на основе стадии и имени
func GetJob(stageName string, jobName string) IPipelineJob {
	jobs := map[string]JobFactory{
		"close.clean":                           NewCleanJob,
		"close.discard-release":                 NewDiscardReleaseJob,
		"close.discard-canary-release":          NewDiscardCanaryReleaseJob,
		"close.close-release":                   NewCloseReleaseJob,
		"close.close-canary-release":            NewCloseCanaryReleaseJob,
		"close.create-tag-from-file":            NewCreateTagFromFileJob,
		"close.delete-merged-branches":          NewDeleteMergedBranchesJob,
		"close.ready-to-prod":                   NewReadyToProdJob,
		"deploy.deploy":                         NewDeployJob,
		"post-deploy.switch":                    NewSwitchJob,
		"pre-build.merge-main":                  NewMergeMainJob,
		"pre-build.validate":                    NewValidationJob,
		"pre-build.sonar":                       NewSonarJob,
		"pre-build.create-badges":               NewCreateBadgesJob,
		"pre-close.create-release-issue":        NewCreateReleaseIssueJob,
		"quality.security-checkov":              NewSecurityCheckJob,
		"quality.security-checkov-dev":          NewSecurityCheckDevJob,
		"quality.security-checkov-stg":          NewSecurityCheckStageJob,
		"quality.sonar":                         NewSonarJob,
		"post-publish.sonar":                    NewSonarJob,
		"post-close.close-jira-release-version": NewCloseJiraReleaseVersionJob,
	}

	factory, ok := jobs[stageName+"."+jobName]
	if !ok {
		slog.Error("Can't find job %s", jobName)
		os.Exit(1)
	}

	job, err := factory()
	if err != nil {
		slog.Error("Can't create job %s: %v", jobName, err)
		os.Exit(1)
	}

	return job
}
