package multimodal

import "github.com/godotask/service"

type MultimodalController struct {
	Service      *service.MultimodalService
	LabelService *service.QuantificationLabelService
}