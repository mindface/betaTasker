package model

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Assessment{},
		&Task{},
		&Memory{},
		&MemoryContext{},
		&Book{},
		&HeuristicsAnalysis{},
	
		&TechnicalFactor{},
		&KnowledgeTransformation{},
		&HeuristicsTracking{},
		&HeuristicsInsight{},
		&HeuristicsPattern{},
		&HeuristicsModel{},
		&MultimodalData{},
		&KnowledgePattern{},
		&LanguageOptimization{},
		&LearningPattern{},
		&PhenomenologicalFramework{},
		&OptimizationModel{},
		&StateEvaluation{},
		&ToolMatchingResult{},

		&ProcessOptimization{},
		&QualitativeLabel{},
		&QuantificationLabel{},
		&TeachingFreeControl{},
		&KnowledgeEntity{},
	}
}
