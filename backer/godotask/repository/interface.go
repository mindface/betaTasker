package repository

import "github.com/godotask/model"

type BookRepositoryInterface interface {
	Create(book *model.Book) error
	FindByID(id string) (*model.Book, error)
	FindAll() ([]model.Book, error)
	Update(id string, book *model.Book) error
	Delete(id string) error
}

type AssessmentRepositoryInterface interface {
	Create(assessment *model.Assessment) error
	FindByID(id string) (*model.Assessment, error)
	FindAll() ([]model.Assessment, error)
	FindByTaskIDAndUserID(userID int, taskID int) ([]model.Assessment, error)
	Update(id string, assessment *model.Assessment) error
	Delete(id string) error
}

type MemoryRepositoryInterface interface {
	Create(memory *model.Memory) error
	FindByID(id string) (*model.Memory, error)
	FindAll() ([]model.Memory, error)
	Update(id string, memory *model.Memory) error
	Delete(id string) error
}

type MemoryContextRepositoryInterface interface {
	FindByCode(code string, contexts *[]model.MemoryContext) error
	FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error
}

type TaskRepositoryInterface interface {
	Create(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	FindAll() ([]model.Task, error)
	Update(id string, task *model.Task) error
	Delete(id string) error
}

type HeuristicsRepositoryInterface interface {
	CreateAnalysis(analysis *model.HeuristicsAnalysis) error
	GetAnalysisById(id string) (*model.HeuristicsAnalysis, error)
	ListAnalyses() ([]model.HeuristicsAnalysis, error)
	UpdateAnalysis(id string, analysis *model.HeuristicsAnalysis) error
	DeleteAnalysis(id string) error
	FindAllAnalyses() ([]model.HeuristicsAnalysis, error)

	CreateTracking(tracking *model.HeuristicsTracking) error
	GetTrackingByUserID(userID string) ([]model.HeuristicsTracking, error)
	DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error)
	CreateModel(model *model.HeuristicsModel) error
}

type HeuristicsInsightRepositoryInterface interface {
	CreateInsight(insight *model.HeuristicsInsight) error
	GetInsightById(id string) (*model.HeuristicsInsight, error)
	ListInsight() ([]model.HeuristicsInsight, error)
	GetInsights(userID string, limit, offset int) ([]model.HeuristicsInsight, int, error)
	UpdateInsight(id string, insight *model.HeuristicsInsight) error
	DeleteInsight(id string) error
}

type HeuristicsPatternRepositoryInterface interface {
	CreatePattern(pattern *model.HeuristicsPattern) error
	GetPatternById(id string) (*model.HeuristicsPattern, error)
	ListPattern() ([]model.HeuristicsPattern, error)
	GetPatterns(userID string, limit, offset int) ([]model.HeuristicsPattern, int, error)
	UpdatePattern(id string, insight *model.HeuristicsPattern) error
	DeletePattern(id string) error
}


type ProcessOptimizationRepositoryInterface interface {
	Create(processOptimization *model.ProcessOptimization) error
	FindByID(id string) (*model.ProcessOptimization, error)
	FindAll() ([]model.ProcessOptimization, error)
	Update(id string, processOptimization *model.ProcessOptimization) error
	Delete(id string) error
}

type PhenomenologicalFrameworkRepositoryInterface interface {
	Create(phenomenologicalFramework *model.PhenomenologicalFramework) error
	FindByID(id string) (*model.PhenomenologicalFramework, error)
	FindAll() ([]model.PhenomenologicalFramework, error)
	Update(id string, phenomenologicalFramework *model.PhenomenologicalFramework) error
	Delete(id string) error
}

type QualitativeLabelRepositoryInterface interface {
	Create(qualitativeLabel *model.QualitativeLabel) error
	FindByID(id string) (*model.QualitativeLabel, error)
	FindAll() ([]model.QualitativeLabel, error)
	Update(id string, qualitativeLabel *model.QualitativeLabel) error
	Delete(id string) error
}

type KnowledgePatternRepositoryInterface interface {
	Create(knowledgePattern *model.KnowledgePattern) error
	FindByID(id string) (*model.KnowledgePattern, error)
	FindAll() ([]model.KnowledgePattern, error)
	Update(id string, knowledgePattern *model.KnowledgePattern) error
	Delete(id string) error
}

type LanguageOptimizationRepositoryInterface interface {
	Create(languageOptimization *model.LanguageOptimization) error
	FindByID(id string) (*model.LanguageOptimization, error)
	FindAll() ([]model.LanguageOptimization, error)
	Update(id string, languageOptimization *model.LanguageOptimization) error
	Delete(id string) error
}

type TeachingFreeControlRepositoryInterface interface {
	Create(teachingFreeControl *model.TeachingFreeControl) error
	FindByID(id string) (*model.TeachingFreeControl, error)
	FindAll() ([]model.TeachingFreeControl, error)
	Update(id string, teachingFreeControl *model.TeachingFreeControl) error
	Delete(id string) error
}
