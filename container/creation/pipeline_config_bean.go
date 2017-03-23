package creation

import (
	"github.com/streamsets/dataextractor/container/common"
	"fmt"
)

type PipelineConfigBean struct {
	Version              string
	ExecutionMode        string
	DeliveryGuarantee    string
	ShouldRetry          bool
	RetryAttempts        float64
	MemoryLimit          string
	MemoryLimitExceeded  string
	NotifyOnStates       []interface{}
	EmailIDs             []interface{}
	Constants            map[string]interface{}
	BadRecordsHandling   string
	StatsAggregatorStage string
	RateLimit            float64
	MaxRunners           float64
}

func NewPipelineConfigBean(pipelineConfig common.PipelineConfiguration) PipelineConfigBean {
	pipelineConfigBean := PipelineConfigBean{}

	for _, config := range pipelineConfig.Configuration {
		switch config.Name {
		case "executionMode":
			pipelineConfigBean.ExecutionMode = config.Value.(string)
			break
		case "deliveryGuarantee":
			pipelineConfigBean.DeliveryGuarantee = config.Value.(string)
			break
		case "shouldRetry":
			pipelineConfigBean.ShouldRetry = config.Value.(bool)
			break
		case "retryAttempts":
			pipelineConfigBean.RetryAttempts = config.Value.(float64)
			break
		case "memoryLimit":
			pipelineConfigBean.MemoryLimit = config.Value.(string)
			break
		case "memoryLimitExceeded":
			pipelineConfigBean.MemoryLimitExceeded = config.Value.(string)
			break
		case "notifyOnStates":
			pipelineConfigBean.NotifyOnStates = config.Value.([]interface{})
			break
		case "emailIDs":
			pipelineConfigBean.EmailIDs = config.Value.([]interface{})
			break
		case "constants":
			constants := config.Value.([]interface{})
			fmt.Println(constants)
			pipelineConfigBean.Constants = make(map[string]interface{})
			/*
			for _, constant := range constants {
				pipelineConfigBean.Constants[constant["key"].(string)] = constant["value"].(interface{})
			}*/
			break
		case "badRecordsHandling":
			pipelineConfigBean.BadRecordsHandling = config.Value.(string)
			break
		case "statsAggregatorStage":
			pipelineConfigBean.StatsAggregatorStage = config.Value.(string)
			break
		case "rateLimit":
			pipelineConfigBean.RateLimit = config.Value.(float64)
			break
		case "maxRunners":
			pipelineConfigBean.MaxRunners = config.Value.(float64)
			break
		}
	}

	return pipelineConfigBean
}