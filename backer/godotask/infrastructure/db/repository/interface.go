package repository

import "github.com/godotask/infrastructure/db/model"

type BookRepositoryInterface interface {
	Create(book *model.Book) error
	FindByID(id string) (*model.Book, error)
	FindAll(userID uint) ([]model.Book, error)
	Update(id string, book *model.Book) error
	Delete(id string) error
}

type AssessmentRepositoryInterface interface {
	Create(assessment *model.Assessment) error
	FindByID(id string) (*model.Assessment, error)
	FindAll(userID uint) ([]model.Assessment, int64, error)
	ListAssessmentsPager(userID uint, offset int, perPage int) ([]model.Assessment, int64, error)
	FindByTaskIDAndUserID(userID int, taskID int) ([]model.Assessment, error)
	ListAssessmentsForTaskUserPager(offset int, perPage int, userID int, taskID int) ([]model.Assessment, int64, error)
	Update(id string, assessment *model.Assessment) error
	Delete(id string) error
}

type MemoryRepositoryInterface interface {
	Create(memory *model.Memory) error
	FindByID(id string) (*model.Memory, error)
	FindAll(userID uint) ([]model.Memory, error)
	ListMemoriesPager(userID uint, offset int, limit int) ([]model.Memory, int64, error)
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
	FindAll(userID uint) ([]model.Task, error)
	ListTasksPager(userID uint, offset int, perPage int) ([]model.Task, int64, error)
	ListTasksByUserPager(userID uint, offset int, perPage int) ([]model.Task, int64, error)
	Update(id string, task *model.Task) error
	Delete(id string) error
}

type HeuristicsRepositoryInterface interface {
	CreateAnalysis(analysis *model.HeuristicsAnalysis) error
	GetAnalysisById(id string) (*model.HeuristicsAnalysis, error)
	ListAnalyses() ([]model.HeuristicsAnalysis, error)
	ListAnalysesPager(userID uint, offset int, limit int) ([]model.HeuristicsAnalysis, int64, error)
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
	ListPattern(userID uint) ([]model.HeuristicsPattern, error)
	GetPatterns(userID string, limit, offset int) ([]model.HeuristicsPattern, int, error)
	UpdatePattern(id string, insight *model.HeuristicsPattern) error
	DeletePattern(id string) error
}

type ProcessOptimizationRepositoryInterface interface {
	Create(processOptimization *model.ProcessOptimization) error
	FindByID(id string) (*model.ProcessOptimization, error)
	FindAll(userID uint) ([]model.ProcessOptimization, error)
	Update(id string, processOptimization *model.ProcessOptimization) error
	Delete(id string) error
}

type PhenomenologicalFrameworkRepositoryInterface interface {
	Create(phenomenologicalFramework *model.PhenomenologicalFramework) error
	FindByID(id string) (*model.PhenomenologicalFramework, error)
	FindAll(userID uint) ([]model.PhenomenologicalFramework, error)
	Update(id string, phenomenologicalFramework *model.PhenomenologicalFramework) error
	Delete(id string) error
}

type QualitativeLabelRepositoryInterface interface {
	Create(qualitativeLabel *model.QualitativeLabel) error
	FindByID(id string) (*model.QualitativeLabel, error)
	FindAll(userID uint) ([]model.QualitativeLabel, error)
	Update(id string, qualitativeLabel *model.QualitativeLabel) error
	Delete(id string) error
}

type KnowledgePatternRepositoryInterface interface {
	Create(knowledgePattern *model.KnowledgePattern) error
	FindByID(id string) (*model.KnowledgePattern, error)
	FindAll(userID uint) ([]model.KnowledgePattern, error)
	Update(id string, knowledgePattern *model.KnowledgePattern) error
	Delete(id string) error
}

type LanguageOptimizationRepositoryInterface interface {
	Create(languageOptimization *model.LanguageOptimization) error
	FindByID(id string) (*model.LanguageOptimization, error)
	FindAll(userID uint) ([]model.LanguageOptimization, error)
	Update(id string, languageOptimization *model.LanguageOptimization) error
	Delete(id string) error
}

type TeachingFreeControlRepositoryInterface interface {
	Create(teachingFreeControl *model.TeachingFreeControl) error
	FindByID(id string) (*model.TeachingFreeControl, error)
	FindAll(userID uint) ([]model.TeachingFreeControl, error)
	Update(id string, teachingFreeControl *model.TeachingFreeControl) error
	Delete(id string) error
}
