package config

import (
	"github.com/caarlos0/env/v10"
	"log/slog"
)

// JobConfig конфиг запуска джоба пайплайна
type JobConfig struct {

	// Debug режим отладки, логируется максимум информации
	Debug bool `env:"DEBUG" envDefault:"false"`

	// JobName содержит имя джоба для запуска
	JobName string `env:"CUSTOM_JOB_NAME,required"`

	// StageName стадия из пайплайна
	StageName string `env:"CI_JOB_STAGE,required"`

	// TrunkMode включен ли режим транк в репозитории
	TrunkMode bool `env:"TRUNK_ENABLE" envDefault:"false"`

	// JiraProjectID ключ проекта в жире
	JiraProjectID string `env:"JIRA_PROJECT_ID"`

	// JiraIssueComponentID ключ компонента в жире
	JiraIssueComponentID string `env:"JIRA_RELEASE_ISSUE_COMPONENT_ID"`

	// JiraWickService url сервиса JiraWick
	JiraWickService string `env:"JIRA_WICK_URL"`

	// JiraNotifyChannel канал для уведомлений о релизах
	JiraNotifyChannel string `env:"JIRA_NOTIFY_CHANNEL" envDefault:"TestNotify"`

	// TechOwners техлиды проекта, используется для оповещения в канале JiraNotifyChannel
	TechOwners string `env:"TECH_OWNERS"`
}

// GetJobConfig фабрика для JobConfig
func GetJobConfig() *JobConfig {
	var appConfig JobConfig
	if err := env.Parse(&appConfig); err != nil {
		slog.Error("Config error: %v", err)
	}
	return &appConfig
}

// GitlabCIVariables переменные окружения, переданные гитлабом
type GitlabCIVariables struct {
	ProjectID     int    `env:"CI_PROJECT_ID,required"`
	ProjectPath   string `env:"CI_PROJECT_PATH"`
	PipelineID    int    `env:"CI_PIPELINE_ID,required"`
	CommitRefName string `env:"CI_COMMIT_REF_NAME,required"`
}

// GetGitlabCIVariables фабрика для GitlabCIVariables
func GetGitlabCIVariables() *GitlabCIVariables {
	var variables GitlabCIVariables
	if err := env.Parse(variables); err != nil {
		slog.Error("Gitlab variables error: %v", err)
	}
	return &variables
}
