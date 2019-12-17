package controller

import (
	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"

	"github.com/giantswarm/master-operator/pkg/project"
)

type MasterConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type Master struct {
	*controller.Controller
}

func NewMaster(config MasterConfig) (*Master, error) {
	var err error

	resourceSets, err := newMasterResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			CRD:          infrastructurev1alpha2.NewMachineDeploymentCRD(),
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(apiv1alpha2.MachineDeployment)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/master-operator-master-controller.
			Name: project.Name() + "-master-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &Master{
		Controller: operatorkitController,
	}

	return c, nil
}

func newMasterResourceSets(config MasterConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := masterResourceSetConfig{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		resourceSet, err = newMasterResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
