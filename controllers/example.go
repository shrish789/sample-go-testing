package controller

import (
	"github.com/shrish789/sample-go/services"
)

type ExampleController struct {
	W services.Service
}

func InitController() *ExampleController {
	return &ExampleController{
		W: services.InitializeWorkflow(),
	}
}