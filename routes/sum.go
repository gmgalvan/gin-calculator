package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-calculator/calculator"
	"github.com/gin-calculator/entities"
	"github.com/gin-gonic/gin"
)

// Result struct is the body result of an operation
type Result struct {
	Result int64 `json:"result"`
}

// Sum godoc
// @Summary      Sum integers
// @Description  sum a list of integers
// @Tags         basicOperations
// @Accept       json
// @Produce      json
// @Param operands body entities.Operands true "operands"
// @Success      200  {object} Result
// @Failure      400  {string} string "error"
// @Router       /sum [post]
func Sum(g *gin.Context) {
	var op entities.Operands
	if err := g.BindJSON(&op); err != nil {
		msgErr := fmt.Sprintf("Error during unmarshalling, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}

	calc := calculator.NewService()
	result := &Result{
		Result: calc.Sum(op),
	}

	g.JSON(http.StatusOK, result)
}
