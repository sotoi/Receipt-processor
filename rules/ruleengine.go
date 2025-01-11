package rules

import "fetch/models"

// Rule is an interface that defines the CalculatePoints method
type RuleRunner interface {
	CalculatePoints(*models.Receipt) int
}

// RuleEngine is a struct that contains a slice of Rule
type RuleEngine struct {
	rules []Rule
}

// NewRuleEngine is a constructor for RuleEngine
func NewRuleEngine(rules []Rule) *RuleEngine {
	return &RuleEngine{
		rules: rules,
	}
}

// CalculatePoints iterates over the rules and applies them to the receipt
func (re *RuleEngine) CalculatePoints(receipt *models.Receipt) int {
	points := 0
	for _, rule := range re.rules {
		points += rule.Apply(receipt)
	}
	return points
}
