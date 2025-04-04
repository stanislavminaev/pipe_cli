package jobs

import "fmt"

type DiscardReleaseJob struct{}

func (j *DiscardReleaseJob) Run() {
	fmt.Println("Running DiscardReleaseJob")
}

func NewDiscardReleaseJob() (IPipelineJob, error) {
	return &DiscardReleaseJob{}, nil
}
