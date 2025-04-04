package jobs

import "fmt"

type DeployJob struct {
}

func (j *DeployJob) Run() {
	fmt.Println("Running DeployJob")
}

func NewDeployJob() (IPipelineJob, error) {
	return &DeployJob{}, nil
}
