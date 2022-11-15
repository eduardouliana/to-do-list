package entity

type TaskRepositoryInterface interface {
	Save(task *Task) error
	Delete(id string) error
	DeleteAll() error
	Read(id string) (*Task, error)
	ReadAll() ([]*Task, error)
}
