package jobs

import "fmt"

type DeleteMergedBranchesJob struct {
}

func (j *DeleteMergedBranchesJob) Run() {
	fmt.Println("Running DeleteMergedBranchesJob")
}

func NewDeleteMergedBranchesJob() (IPipelineJob, error) {
	return &DeleteMergedBranchesJob{}, nil
}
