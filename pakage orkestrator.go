package orchestrator

import (
	"calc2/internal/evaluator"
	"calc2/internal/parser"
	"fmt"
)

type Orchestrator struct {
	parser    *parser.Parser
	evaluator *evaluator.Evaluator
}

func New() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) Calculate(expression string) (float64, error) {
	tokens, err := parser.Parse(expression)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %v", err)
	}

	result, err := evaluator.Evaluate(tokens)
	if err != nil {
		return 0, fmt.Errorf("evaluation error: %v", err)
	}

	return result, nil
}

func (o *Orchestrator) Validate(expression string) error {
	_, err := parser.Parse(expression)
	return err
}
