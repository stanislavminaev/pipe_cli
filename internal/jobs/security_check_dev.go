package jobs

import "fmt"

type SecurityCheckDevJob struct {
}

func (j *SecurityCheckDevJob) Run() {
	fmt.Println("Running SecurityCheckDevJob")
}

func NewSecurityCheckDevJob() (IPipelineJob, error) {
	return &SecurityCheckDevJob{}, nil
}
