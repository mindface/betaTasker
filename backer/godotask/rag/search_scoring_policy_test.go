package rag

import (
	"math"
	"testing"
)

func TestSearchScoringPolicyScore(t *testing.T) {
	policy := MedicalSearchScoringPolicy()
	input := SearchScoreInput{
		DocumentID:      "doc:001",
		TextSimilarity:  0.72,
		OntologyMatch:   1.00,
		RiskPriority:    0.85,
		GraphCentrality: 0.40,
	}

	score, err := policy.Score(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 0.825
	if math.Abs(score-want) > 0.000001 {
		t.Fatalf("score = %v, want %v", score, want)
	}
}

func TestSearchScoringPolicyNormalize(t *testing.T) {
	policy := SearchScoringPolicy{
		Alpha: 2,
		Beta:  2,
		Gamma: 1,
		Delta: 0,
	}

	normalized, err := policy.Normalize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if normalized.Alpha != 0.4 || normalized.Beta != 0.4 || normalized.Gamma != 0.2 || normalized.Delta != 0 {
		t.Fatalf("unexpected normalized policy: %+v", normalized)
	}
}

func TestSearchScoringPolicyRejectsInvalidWeights(t *testing.T) {
	tests := []struct {
		name   string
		policy SearchScoringPolicy
	}{
		{
			name:   "negative",
			policy: SearchScoringPolicy{Alpha: -1, Beta: 1},
		},
		{
			name:   "zero total",
			policy: SearchScoringPolicy{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.policy.Validate(); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestSearchScoringPolicyRank(t *testing.T) {
	policy := MedicalSearchScoringPolicy()
	inputs := []SearchScoreInput{
		{
			DocumentID:      "text-only",
			TextSimilarity:  0.95,
			OntologyMatch:   0.10,
			RiskPriority:    0.10,
			GraphCentrality: 0.20,
		},
		{
			DocumentID:      "ontology-risk",
			TextSimilarity:  0.70,
			OntologyMatch:   1.00,
			RiskPriority:    0.90,
			GraphCentrality: 0.40,
		},
	}

	results, err := policy.Rank(inputs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := results[0].DocumentID; got != "ontology-risk" {
		t.Fatalf("top document = %s, want ontology-risk", got)
	}
}

func TestSearchScoreInputClamp(t *testing.T) {
	input := SearchScoreInput{
		TextSimilarity:  -0.2,
		OntologyMatch:   1.2,
		RiskPriority:    0.5,
		GraphCentrality: 2.0,
	}.Clamp()

	if input.TextSimilarity != 0 || input.OntologyMatch != 1 || input.RiskPriority != 0.5 || input.GraphCentrality != 1 {
		t.Fatalf("unexpected clamped input: %+v", input)
	}
}
