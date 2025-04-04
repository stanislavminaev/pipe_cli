package jobs

import "fmt"

type CreateBadgesJob struct {
}

func (j *CreateBadgesJob) Run() {
	fmt.Println("Running CreateBadgesJob")
}

func NewCreateBadgesJob() (IPipelineJob, error) {
	return &CreateBadgesJob{}, nil
}
