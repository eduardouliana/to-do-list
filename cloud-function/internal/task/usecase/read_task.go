package usecase

import (
	"github.com/eduardouliana/to-do-list/internal/task/entity"
	"github.com/eduardouliana/to-do-list/internal/task/infra/database"
)

type ReadTaskInputDTO struct {
	Id string
}

type ReadTaskOutputDTO struct {
	Id          string
	Description string
	Done        bool
}

type ReadTaskUseCase struct {
	TaskRepository entity.TaskRepositoryInterface
}

func NewReadTaskUseCase(taskRepository database.TaskRepository) *ReadTaskUseCase {
	return &ReadTaskUseCase{
		TaskRepository: &taskRepository,
	}
}

func (s *ReadTaskUseCase) Execute(input ReadTaskInputDTO) (*ReadTaskOutputDTO, error) {
	task, err := s.TaskRepository.Read(input.Id)
	if err != nil {
		return nil, err
	}

	return &ReadTaskOutputDTO{
		Id:          task.Id,
		Description: task.Description,
		Done:        task.Done,
	}, nil
}

func (s *ReadTaskUseCase) ExecuteAll() ([]*ReadTaskOutputDTO, error) {
	taskList, err := s.TaskRepository.ReadAll()
	if err != nil {
		return nil, err
	}

	outputDTOList := []*ReadTaskOutputDTO{}

	for _, task := range taskList {
		outputDTO := ReadTaskOutputDTO{
			Id:          task.Id,
			Description: task.Description,
			Done:        task.Done,
		}
		outputDTOList = append(outputDTOList, &outputDTO)
	}

	return outputDTOList, nil
}
