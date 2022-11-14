package reader

import "github.com/inspectorvitya/linxdatacenter-test-task/internal/model"

type Reader interface {
	Read() (model.Product, error)
}
