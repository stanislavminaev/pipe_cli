package jobs

import "fmt"

type SecurityCheckJob struct {
}

func (j *SecurityCheckJob) Run() {
	fmt.Println("Running SecurityCheckJob")
}

func NewSecurityCheckJob() (IPipelineJob, error) {
	return &SecurityCheckJob{}, nil
}
