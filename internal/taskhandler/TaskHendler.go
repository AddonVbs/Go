package taskhandler

import (
	ts "BackEnd/internal/taskservice"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service ts.TaskServers
}

func NewTaskHandler(s ts.TaskServers) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	task, err := h.service.GetAllTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ts.Response{Status: "error", Message: "Cloud not get task "})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid payload"})
	}
	task, err := h.service.CreateTask(req.Expression)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ts.Response{Status: "error", Message: "Could not create task"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	var req struct {
		Expression string `json:"expression"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid req"})
	}

	updata, err := h.service.UpdataTask(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid updata"})
	}

	return c.JSON(http.StatusOK, updata)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid ID"})
	}

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusBadRequest, ts.Response{Status: "error", Message: "Invalid del."})
	}

	return c.JSON(http.StatusNoContent, ts.Response{Status: "Seccess", Message: "Was Del."})
}
