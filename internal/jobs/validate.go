package jobs

import "fmt"

// ValidationJob — структура, реализующая джоб валидации
type ValidationJob struct{}

// Run - запуск джоба
func (j *ValidationJob) Run() {
	fmt.Println("Running ValidationJob")
}

// NewValidationJob — конструктор, возвращающий PipelineJob
func NewValidationJob() (IPipelineJob, error) {
	return &ValidationJob{}, nil
}
