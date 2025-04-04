package jobs

import "fmt"

type SwitchJob struct {
}

func (j *SwitchJob) Run() {
	fmt.Println("Running SwitchJob")
}

func NewSwitchJob() (IPipelineJob, error) {
	return &SwitchJob{}, nil
}
