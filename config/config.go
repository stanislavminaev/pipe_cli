package config

import (
	"github.com/caarlos0/env/v10"
	"log/slog"
)

type JobConfig struct {
	Debug                bool   `env:"DEBUG" envDefault:"false"`
	JobName              string `env:"CUSTOM_JOB_NAME,required"`
	StageName            string `env:"CI_JOB_STAGE,required"`
	TrunkMode            bool   `env:"TRUNK_ENABLE" envDefault:"false"`
	JiraProjectID        string `env:"JIRA_PROJECT_ID"`
	JiraIssueComponentID string `env:"JIRA_RELEASE_ISSUE_COMPONENT_ID"`
	JiraWickService      string `env:"JIRA_WICK_URL"`
	JiraNotifyChannel    string `env:"JIRA_NOTIFY_CHANNEL" envDefault:"TestNotify"`
	TechOwners           string `env:"TECH_OWNERS"`
}

func GetJobConfig() *JobConfig {
	var appConfig JobConfig
	if err := env.Parse(&appConfig); err != nil {
		slog.Error("Config error: %v", err)
	}
	return &appConfig
}

type GitlabCIVariables struct {
	ProjectID     int    `env:"CI_PROJECT_ID,required"`
	ProjectPath   string `env:"CI_PROJECT_PATH"`
	PipelineID    int    `env:"CI_PIPELINE_ID,required"`
	CommitRefName string `env:"CI_COMMIT_REF_NAME,required"`
}

func GetGitlabCIVariables() *GitlabCIVariables {
	var variables GitlabCIVariables
	if err := env.Parse(variables); err != nil {
		slog.Error("Gitlab variables error: %v", err)
	}
	return &variables
}
