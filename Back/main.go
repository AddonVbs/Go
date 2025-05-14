package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var task string = ""

var calculations = []Calculation{}

var Task string = ""

type task struct {
	Tasks string `json:"task"`
}

func PostHendler(t task) string {
	Task = t.Tasks
	return Task

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

}
func postCalculation(c echo.Context) error {
	var req CalculationRequest
	err := c.Bind(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invaild request"})
	}

	result, err := calculationExpression(req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid Expression"})
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}
	calculations = append(calculations, calc)

	return c.JSON(http.StatusCreated, calc)

}

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)

	e.POST("/calculations", postCalculation)
	e.Start("localhost:8080")

}
