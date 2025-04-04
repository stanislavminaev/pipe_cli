package jobs

import "fmt"

type CloseCanaryReleaseJob struct {
}

func (j *CloseCanaryReleaseJob) Run() {
	fmt.Println("Running CloseCanaryReleaseJob")
}

func NewCloseCanaryReleaseJob() (IPipelineJob, error) {
	return &CloseCanaryReleaseJob{}, nil
}
