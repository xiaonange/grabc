package grbac

import (
	"fmt"
)

//用户接口
type Identify interface {
	GetUid() int
}

type Rbac struct {
}

func (this *Rbac) CheckAccess() bool {
	return false
}