package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type StaticWorkerPool struct {
	WorkerPool
}

// NewStaticWorkerPool creates and initializes a static worker pool.
func NewStaticWorkerPool(name string) *StaticWorkerPool {
	return &StaticWorkerPool{
		WorkerPool: *newWorkerPool(name, WorkerPoolTypeStatic),
	}
}

func (s *StaticWorkerPool) GetIsDefault() bool {
	return s.IsDefault
}

// Validate checks the state of the static worker pool and returns an error if
// invalid.
func (s *StaticWorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(s)
}
