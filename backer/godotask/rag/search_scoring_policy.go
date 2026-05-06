package rag

import (
	"errors"
	"sort"
)

// SearchScoringPolicy defines how hybrid search candidates are ranked.
// TextSimilarity is usually produced by vector search, while the other values
// can be retrieved from an RDF store through SPARQL.
type SearchScoringPolicy struct {
	Alpha float64 `json:"alpha"` // semantic text similarity
	Beta  float64 `json:"beta"`  // ontology concept match
	Gamma float64 `json:"gamma"` // risk or urgency priority
	Delta float64 `json:"delta"` // ontology graph centrality
}

// SearchScoreInput is the normalized scoring input for one candidate document.
// Each field should be in the 0.0 to 1.0 range before scoring.
type SearchScoreInput struct {
	DocumentID       string  `json:"document_id"`
	TextSimilarity   float64 `json:"text_similarity"`
	OntologyMatch    float64 `json:"ontology_match"`
	RiskPriority     float64 `json:"risk_priority"`
	GraphCentrality  float64 `json:"graph_centrality"`
}

type SearchScoreResult struct {
	DocumentID string           `json:"document_id"`
	Score      float64          `json:"score"`
	Input      SearchScoreInput `json:"input"`
}

func DefaultSearchScoringPolicy() SearchScoringPolicy {
	return SearchScoringPolicy{
		Alpha: 0.40,
		Beta:  0.30,
		Gamma: 0.20,
		Delta: 0.10,
	}
}

func MedicalSearchScoringPolicy() SearchScoringPolicy {
	return SearchScoringPolicy{
		Alpha: 0.25,
		Beta:  0.35,
		Gamma: 0.30,
		Delta: 0.10,
	}
}

func (p SearchScoringPolicy) Validate() error {
	if p.Alpha < 0 || p.Beta < 0 || p.Gamma < 0 || p.Delta < 0 {
		return errors.New("search scoring policy weights must be non-negative")
	}
	if p.TotalWeight() == 0 {
		return errors.New("search scoring policy must have at least one positive weight")
	}
	return nil
}

func (p SearchScoringPolicy) TotalWeight() float64 {
	return p.Alpha + p.Beta + p.Gamma + p.Delta
}

func (p SearchScoringPolicy) Normalize() (SearchScoringPolicy, error) {
	if err := p.Validate(); err != nil {
		return SearchScoringPolicy{}, err
	}

	total := p.TotalWeight()
	return SearchScoringPolicy{
		Alpha: p.Alpha / total,
		Beta:  p.Beta / total,
		Gamma: p.Gamma / total,
		Delta: p.Delta / total,
	}, nil
}

func (p SearchScoringPolicy) Score(input SearchScoreInput) (float64, error) {
	normalizedPolicy, err := p.Normalize()
	if err != nil {
		return 0, err
	}

	input = input.Clamp()
	return normalizedPolicy.Alpha*input.TextSimilarity +
		normalizedPolicy.Beta*input.OntologyMatch +
		normalizedPolicy.Gamma*input.RiskPriority +
		normalizedPolicy.Delta*input.GraphCentrality, nil
}

func (p SearchScoringPolicy) Rank(inputs []SearchScoreInput) ([]SearchScoreResult, error) {
	results := make([]SearchScoreResult, 0, len(inputs))
	for _, input := range inputs {
		score, err := p.Score(input)
		if err != nil {
			return nil, err
		}
		results = append(results, SearchScoreResult{
			DocumentID: input.DocumentID,
			Score:      score,
			Input:      input.Clamp(),
		})
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results, nil
}

func (i SearchScoreInput) Clamp() SearchScoreInput {
	i.TextSimilarity = clamp01(i.TextSimilarity)
	i.OntologyMatch = clamp01(i.OntologyMatch)
	i.RiskPriority = clamp01(i.RiskPriority)
	i.GraphCentrality = clamp01(i.GraphCentrality)
	return i
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
