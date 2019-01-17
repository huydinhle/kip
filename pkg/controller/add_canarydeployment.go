package controller

import (
	"github.com/huydinhle/kip/pkg/controller/canarydeployment"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, canarydeployment.Add)
}