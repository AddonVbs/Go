package taskhendler

import (
	ts "BackEnd/internal/TaskService"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskHendler struct {
	service ts.TaskServers
}

func NewTaskHendler(s ts.TaskServers) *TaskHendler {
	return &TaskHendler{service: s}

}

func (h *TaskHendler) GetHandler(c echo.Context) error {
	task, err := h.service.GetAllTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ts.Response{Status: "error", Message: "Cloud not get task "})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHendler) PostHandler(c echo.Context) error {
	var req ts.RepositorysTasks

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid payload"})
	}

	task, err := h.service.CreateTask(req.expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid req"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHendler) PatchHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	var req ts.TaskRepository
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid req"})
	}

	updata, err := h.service.UpdataTask(id, req.expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid updata"})
	}

	return c.JSON(http.StatusOK, updata)
}

func (h *TaskHendler) DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid del."})
	}

	return c.JSON(http.StatusNoContent, ts.Response{Status: "Seccess", Message: "Was Del."})
}
