package main

import (
	"os"
	"pipe-cli/config"
	"pipe-cli/internal/jobs"
)

func main() {
	dev()
	cfg := config.GetJobConfig()
	job := jobs.GetJob(cfg.StageName, cfg.JobName)
	job.Run()
}

func dev() {
	//os.Setenv("CI_JOB_STAGE", "pre-build")
	//os.Setenv("CUSTOM_JOB_NAME", "validate")
	os.Setenv("CI_JOB_STAGE", "post-close")
	os.Setenv("CUSTOM_JOB_NAME", "close-jira-release-version")

}
