package jobs

import "fmt"

type CreateReleaseIssueJob struct {
}

func (j *CreateReleaseIssueJob) Run() {
	fmt.Println("Running CreateReleaseIssueJob")
}

func NewCreateReleaseIssueJob() (IPipelineJob, error) {
	return &CreateReleaseIssueJob{}, nil
}
