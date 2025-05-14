package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type task struct {
	Tasks string `json:"task"`
}

<<<<<<< HEAD
var Task string = "Содить за молоком"
=======
func PostHendler(t task) string {
	Task = t.Tasks
	return Task
>>>>>>> aa5a94d1da5cf41888215bc56ccaca7eb10ac6ac

// 1
func PostHendler(h echo.Context) error {

<<<<<<< HEAD
	return h.JSON(http.StatusOK, map[string]string{"task": Task})
=======
func helloTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello-")

}

func calculationExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
>>>>>>> aa5a94d1da5cf41888215bc56ccaca7eb10ac6ac

}

func main() {

	PostHendler(task{Tasks: "уборка"})
	e := echo.New()

	http.HandleFunc("/", helloTask)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/myTask", PostHendler)

	e.Start("localhost:8080")

}
