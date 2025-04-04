package jobs

import "fmt"

type SonarJob struct {
}

func (j *SonarJob) Run() {
	fmt.Println("Running SonarJob")
}

func NewSonarJob() (IPipelineJob, error) {
	return &SwitchJob{}, nil
}
