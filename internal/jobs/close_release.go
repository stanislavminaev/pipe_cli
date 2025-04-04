package jobs

import "fmt"

type CloseReleaseJob struct {
}

func (j *CloseReleaseJob) Run() {
	fmt.Println("Running CloseReleaseJob")
}

func NewCloseReleaseJob() (IPipelineJob, error) {
	return &CloseReleaseJob{}, nil
}
