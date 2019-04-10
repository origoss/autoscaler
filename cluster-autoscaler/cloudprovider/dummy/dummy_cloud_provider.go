package dummy

import (
	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	kubeclient "k8s.io/client-go/kubernetes"
)

const (
	ProviderName = "dummy"
)

type DummyCloudProvider struct {
	*kubeclient.Clientset
}

func (dcp *DummyCloudProvider) Name() string {
	glog.Info("DummyCloudProvider.Name called")
	return "dummy"
}

func (dcp *DummyCloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	return []cloudprovider.NodeGroup{NodeGroup}
}

func (dcp *DummyCloudProvider) NodeGroupForNode(*apiv1.Node) (cloudprovider.NodeGroup, error) {
	return NodeGroup, nil
}

func (dcp *DummyCloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	return nil, nil
}

func (dcp *DummyCloudProvider) GetAvailableMachineTypes() ([]string, error) {
	return nil, nil
}

func (dcp *DummyCloudProvider) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string, taints []apiv1.Taint, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	return nil, nil
}

func (dcp *DummyCloudProvider) GetResourceLimiter() (*cloudprovider.ResourceLimiter, error) {
	glog.Info("DummyCloudProvider.GetResourceLimiter called")
	return DummyResourceLimiter, nil
}

func (dcp *DummyCloudProvider) GetInstanceID(node *apiv1.Node) string {
	return string(node.Name)
}

func (dcp *DummyCloudProvider) Cleanup() error {
	glog.Info("DummyCloudProvider.Cleanup called")
	return nil
}

func (dcp *DummyCloudProvider) Refresh() error {
	glog.Info("DummyCloudProvider.Refresh called")
	nodeList, err := dcp.Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		glog.Error("Cannot list nodes: %+v", err)
	}
	NodeGroup.nodes = nodeList.Items
	return nil
}
