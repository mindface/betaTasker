package model

import (
	"github.com/godotask/lib"
)

type JSON = lib.JSON

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
		&KnowledgePattern{},
		&LanguageOptimization{},
		&ProcessOptimization{},
		&QualitativeLabel{},
		&QuantificationLabel{},
		&TeachingFreeControl{},
		&KnowledgeEntity{},
	}
}
