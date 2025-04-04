package jobs

import "fmt"

type CreateTagFromFileJob struct {
}

func (j *CreateTagFromFileJob) Run() {
	fmt.Println("Running CreateTagFromFileJob")
}

func NewCreateTagFromFileJob() (IPipelineJob, error) {
	return &CreateTagFromFileJob{}, nil
}
