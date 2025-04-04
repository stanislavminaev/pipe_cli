package jobs

import "fmt"

type ReadyToProdJob struct {
}

func (j *ReadyToProdJob) Run() {
	fmt.Println("Running ReadyToProdJob")
}

func NewReadyToProdJob() (IPipelineJob, error) {
	return &ReadyToProdJob{}, nil
}
