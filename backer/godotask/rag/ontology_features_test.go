package rag

import "testing"

func TestBuildScoreInputUsesExactRelatedRiskAndCentrality(t *testing.T) {
	registry := NewOntologyRegistry()
	registry.RegisterMany([]OntologyConcept{
		{
			ID:         "HP:0001945",
			Label:      "fever",
			ParentIDs:  []string{"HP:SYM:000"},
			RelatedIDs: []string{"DX:001", "OB:T:001"},
			Metadata:   map[string]interface{}{"risk_level": 3, "urgency": 3},
			Domain:     "medical",
		},
		{
			ID:         "DX:001",
			Label:      "pneumonia",
			ParentIDs:  []string{"DX:000"},
			RelatedIDs: []string{"HP:0001945", "HP:0002090"},
			Metadata:   map[string]interface{}{"risk_level": 4, "criticality": 4},
			Domain:     "medical",
		},
		{
			ID:         "HP:0002090",
			Label:      "dyspnea",
			ParentIDs:  []string{"HP:SYM:000"},
			RelatedIDs: []string{"DX:001"},
			Metadata:   map[string]interface{}{"risk_level": 5, "urgency": 5},
			Domain:     "medical",
		},
	})
	registry.BuildFeatures()

	input := registry.BuildScoreInput(
		[]string{"HP:0001945", "DX:001"},
		0.72,
		OntologyDocument{
			ID:           "DOC:MED:001",
			OntologyTags: []string{"HP:0001945", "DX:001", "HP:0002090"},
			Metadata:     map[string]interface{}{"risk_level": 4},
		},
	)

	if input.OntologyMatch <= 0.5 {
		t.Fatalf("ontology match too low: %+v", input)
	}
	if input.RiskPriority < 0.8 {
		t.Fatalf("risk priority too low: %+v", input)
	}
	if input.GraphCentrality <= 0 {
		t.Fatalf("graph centrality should be derived: %+v", input)
	}
}

func TestRankOntologyDocumentsPrefersIndustrialEmergencyProcedure(t *testing.T) {
	registry := NewOntologyRegistry()
	registry.RegisterMany([]OntologyConcept{
		{
			ID:         "EQ:002",
			Label:      "hydraulic press",
			ParentIDs:  []string{"EQ:000"},
			RelatedIDs: []string{"OB:002", "AN:002", "MT:002", "SP:001"},
			Metadata:   map[string]interface{}{"risk_level": 5, "criticality": 5},
			Domain:     "industrial",
		},
		{
			ID:         "AN:002",
			Label:      "hydraulic leak",
			ParentIDs:  []string{"AN:000"},
			RelatedIDs: []string{"EQ:002", "OB:002", "SP:001"},
			Metadata:   map[string]interface{}{"risk_level": 5, "criticality": 5},
			Domain:     "industrial",
		},
		{
			ID:         "SP:001",
			Label:      "lockout tagout",
			ParentIDs:  []string{"SP:000"},
			RelatedIDs: []string{"EQ:002", "AN:002"},
			Metadata:   map[string]interface{}{"risk_level": 5, "criticality": 5},
			Domain:     "industrial",
		},
	})
	registry.BuildFeatures()

	docs := []OntologyDocument{
		{
			ID:           "procedure",
			OntologyTags: []string{"EQ:002", "AN:002", "SP:001"},
			Metadata:     map[string]interface{}{"risk_level": 5},
		},
		{
			ID:           "text-only",
			OntologyTags: []string{"EQ:002"},
			Metadata:     map[string]interface{}{"risk_level": 2},
		},
	}

	results, err := RankOntologyDocuments(
		registry,
		SearchScoringPolicy{Alpha: 0.25, Beta: 0.40, Gamma: 0.25, Delta: 0.10},
		[]string{"EQ:002", "AN:002"},
		map[string]float64{
			"procedure": 0.65,
			"text-only": 0.95,
		},
		docs,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := results[0].DocumentID; got != "procedure" {
		t.Fatalf("top document = %s, want procedure", got)
	}
}
