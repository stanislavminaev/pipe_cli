package jobs

import "fmt"

type CleanJob struct{}

func (j *CleanJob) Run() {
	fmt.Println("Running CleanJob")
}

func NewCleanJob() (IPipelineJob, error) {
	return &CleanJob{}, nil
}
