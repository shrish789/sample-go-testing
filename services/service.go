package services

import (
	"github.com/shrish789/sample-go/dao"
)

type Service interface {
	Example() string
}

type Workflow struct {
	Db    dao.DataLayer
}

func InitializeWorkflow() Service {
	return &Workflow{
		Db: dao.InitializeDao(),
	}
}