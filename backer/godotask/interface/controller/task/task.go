package task

import "github.com/godotask/usecase/service"

type TaskController struct {
	Service *service.TaskService
	KnowledgeEntityService *service.KnowledgeEntityService
}
