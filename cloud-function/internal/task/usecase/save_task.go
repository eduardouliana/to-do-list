package usecase

import (
	"github.com/eduardouliana/to-do-list/internal/task/entity"
	"github.com/eduardouliana/to-do-list/internal/task/infra/database"
)

type SaveTaskInputDTO struct {
	Id          string
	Description string
	Done        bool
}

type SaveTaskOutputDTO struct {
	Status bool
}

type SaveTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewSaveTaskUseCase(taskRepository database.TaskRepository) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		TaskRepository: &taskRepository,
	}
}

func (s *SaveTaskUseCase) Execute(input SaveTaskInputDTO) (*SaveTaskOutputDTO, error) {
	task, err := entity.NewTask(input.Id, input.Description, input.Done)
	if err != nil {
		return nil, err
	}

	err = s.TaskRepository.Save(task)
	if err != nil {
		return nil, err
	}

	return &SaveTaskOutputDTO{
		Status: true,
	}, nil
}
