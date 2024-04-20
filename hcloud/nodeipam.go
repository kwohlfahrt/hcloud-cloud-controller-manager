package hcloud

import (
	"context"

	cp "k8s.io/cloud-provider"
	cpapp "k8s.io/cloud-provider/app"
	cpconfig "k8s.io/cloud-provider/app/config"
	cmapp "k8s.io/controller-manager/app"
	cmcontroller "k8s.io/controller-manager/controller"
)

type NodeIpamController struct {
}

func NodeIpamControllerConstructor(initcontext cpapp.ControllerInitContext, completedConfig *cpconfig.CompletedConfig, cloud cp.Interface) cpapp.InitFunc {
	return func(ctx context.Context, controllerContext cmapp.ControllerContext) (controller cmcontroller.Interface, enabled bool, err error) {
		return nil, true, nil
	}
}
