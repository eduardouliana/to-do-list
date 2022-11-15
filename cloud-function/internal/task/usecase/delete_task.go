package usecase

import (
	"github.com/eduardouliana/to-do-list/internal/task/entity"
	"github.com/eduardouliana/to-do-list/internal/task/infra/database"
)

type DeleteTaskInputDTO struct {
	Id string
}

type DeleteTaskOutputDTO struct {
	Status bool
}

type DeleteTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewDeleteTaskUseCase(taskRepository database.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		TaskRepository: &taskRepository,
	}
}

func (s *DeleteTaskUseCase) Execute(input DeleteTaskInputDTO) (*DeleteTaskOutputDTO, error) {
	err := s.TaskRepository.Delete(input.Id)
	if err != nil {
		return nil, err
	}

	return &DeleteTaskOutputDTO{
		Status: true,
	}, nil
}

func (s *DeleteTaskUseCase) ExecuteAll() error {
	err := s.TaskRepository.DeleteAll()
	if err != nil {
		return err
	}

	return nil
}
