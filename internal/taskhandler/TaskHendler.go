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
	return h.JSON(http.StatusOK, task)
}

func (h *StrictTaskHandler) PostHandler(c echo.Context) error {
	var req struct {
		Task string `json:"task"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid payload"})
	}
	task, err := h.service.CreateTask(req.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ts.Response{Status: "error", Message: "Could not create task"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *StrictTaskHandler) PatchHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
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
