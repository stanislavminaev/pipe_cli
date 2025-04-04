package jobs

import "fmt"

type DiscardCanaryReleaseJob struct {
}

func (j *DiscardCanaryReleaseJob) Run() {
	fmt.Println("Running DiscardCanaryReleaseJob")
}

func NewDiscardCanaryReleaseJob() (IPipelineJob, error) {
	return &DiscardCanaryReleaseJob{}, nil
}
