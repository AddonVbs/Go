package taskservers

type TaskServers interface {
	CreateTask(expression string) (Task, error)
	GetAllTask() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	UpdataTask(id int, expression string) (Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) taskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(expression string) (Task, error) {
	var id int
	t := Task{
		ID:    id,
		Task1: expression,
	}

	if err := s.repo.CreateTask(t); err != nil {
		return Task{}, err
	}
	id++
	return t, nil

}

func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()

}

func (s *taskService) GetTaskByID(id int) (Task, error) {
	return s.repo.GetTaskByID(id)

}

func (s *taskService) UpdataTask(id int, expression string) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Task1 = expression

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err

	}
	return task, nil

}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
