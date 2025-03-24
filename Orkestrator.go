package orchestrator

import (
	"calc2/internal/evaluator"
	"calc2/internal/parser"
	"fmt"
)

// Orchestrator управляет процессом вычисления выражения
type Orchestrator struct {
	parser    *parser.Parser
	evaluator *evaluator.Evaluator
}

// New создает новый экземпляр Orchestrator
func New() *Orchestrator {
	return &Orchestrator{}
}

// Calculate вычисляет результат математического выражения
func (o *Orchestrator) Calculate(expression string) (float64, error) {
	// Шаг 1: Парсинг выражения в токены
	tokens, err := parser.Parse(expression)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %v", err)
	}

	// Шаг 2: Вычисление выражения
	result, err := evaluator.Evaluate(tokens)
	if err != nil {
		return 0, fmt.Errorf("evaluation error: %v", err)
	}

	return result, nil
}

// Validate проверяет выражение на корректность
func (o *Orchestrator) Validate(expression string) error {
	_, err := parser.Parse(expression)
	return err
}p
