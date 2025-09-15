package model

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Assessment{},
		&Task{},
		&Memory{},
		&MemoryContext{},
		&TechnicalFactor{},
		&KnowledgeTransformation{},
		&Book{},
		&HeuristicsAnalysis{},
		&HeuristicsTracking{},
		&HeuristicsInsight{},
		&HeuristicsPattern{},
		&HeuristicsModel{},
		&MultimodalData{},
		&KnowledgePattern{},
		&LanguageOptimization{},
		&ProcessOptimization{},
		&QualitativeLabel{},
		&QuantificationLabel{},
		&TeachingFreeControl{},
	}
}
