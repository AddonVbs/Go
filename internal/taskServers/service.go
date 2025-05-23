package taskservers

type TaskServers interface {
	CreateTask(expression string) (Task, error)
	GetAllTask() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdataTask(id, expression string) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) taskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(expression string) (Task, error) {

}

func (s *taskService) GetAllTask() ([]Task, error) {

}

func (s *taskService) GetTaskByID(id string) (Task, error) {

}

func (s *taskService) UpdataTask(id, expression string) (Task, error) {

}

func (s *taskService) DeleteTask(id string) error {

}
