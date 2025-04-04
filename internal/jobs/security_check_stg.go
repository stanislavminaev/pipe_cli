package jobs

import "fmt"

type SecurityCheckStageJob struct {
}

func (j *SecurityCheckStageJob) Run() {
	fmt.Println("Running SecurityCheckStageJob")
}

func NewSecurityCheckStageJob() (IPipelineJob, error) {
	return &SecurityCheckStageJob{}, nil
}
