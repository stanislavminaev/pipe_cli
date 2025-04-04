package jobs

import "fmt"

type MergeMainJob struct {
}

func (j *MergeMainJob) Run() {
	fmt.Println("Running MergeMainJob")
}

func NewMergeMainJob() (IPipelineJob, error) {
	return &MergeMainJob{}, nil
}
