package calculator

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result, err := Calculate(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}

func Calculate(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	valid, _ := regexp.MatchString(`^[\d\+\-\*/().]+$`, expression)
	if !valid {
		return 0, ErrInvalidExpression
	}

	if strings.Contains(expression, "/0") {
		return 0, ErrDivisionByZero
	}

	return eval(expression)
}

func eval(expression string) (float64, error) {
	// Реализация парсинга и вычислений
	return strconv.ParseFloat(expression, 64) // Заглушка
}
