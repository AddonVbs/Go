package taskhandler

import (
	"context"
	"net/http"
	"strconv"

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

func (h *StrictTaskHandler) GetHandler(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
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

func (h *StrictTaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
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

func (h *StrictTaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	pathParams := request.PathParams
	body := request.Body

	id := int(*pathParams.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	var req struct {
		Task string `json:"task"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid req"})
	}

	updata, err := h.service.UpdataTask(id, req.Task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid updata"})
	}

	return c.JSON(http.StatusOK, updata)
}

func (h *StrictTaskHandler) DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid del."})
	}

	return c.JSON(http.StatusNoContent, ts.Response{Status: "Seccess", Message: "Was Del."})
}
