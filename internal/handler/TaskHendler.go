package handler

import (
	"context"
	"net/http"

	"BackEnd/internal/taskservice"
	"BackEnd/internal/web/tasks"

	"github.com/labstack/echo/v4"
)

type StrictTaskHandler struct {
	service taskservice.TaskServers
}

func NewStrictTaskHandler(s taskservice.TaskServers) *StrictTaskHandler {
	return &StrictTaskHandler{service: s}
}

func (h *StrictTaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	task, err := h.service.GetAllTask()

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, srvTask := range task {

		var idint int = int(srvTask.ID)
		var textPtr = srvTask.Task
		response = append(response, tasks.Task{
			Id:   &idint,
			Task: &textPtr,
		})
	}

	return response, nil
}

func (h *StrictTaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body

	newServiceTask := taskservice.Task{
		Task: *body.Task,
	}

	task, err := h.service.CreateTask(newServiceTask.Task)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var idint int = int(task.ID)
	var textPtr = task.Task
	resp := tasks.PostTasks201JSONResponse{
		Id:   &idint,
		Task: &textPtr,
		// IsDone: nil,
	}

	return resp, nil
}

func (h *StrictTaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := int(request.Id)

	if err := h.service.DeleteTask(id); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *StrictTaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	pathParams := request.Id
	body := request.Body

	id := int(pathParams)

	updata, err := h.service.UpdataTask(id, *body.Task)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var idPach int = (updata.ID)
	bodyPach := updata.Task
	resp := tasks.PatchTasksId200JSONResponse{
		Id:   &idPach,
		Task: &bodyPach,
	}

	return resp, nil
}

func (h *StrictTaskHandler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksByUserIDRequestObject) (tasks.GetTasksByUserIDResponseObject, error) {
	userID := request.UserId
	ts, err := h.service.GetTasksForUser(userID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return tasks.GetTasksByUserID200JSONResponse(ts), nil
}
