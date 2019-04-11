package dummy

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	"net/http"
)

var dummyWebHook = flag.String("dummyhook", "", "An URL where a PUT request is sent whenever the ClusterAutoscaler detects the need for scaling the cluster. ")

const (
	NodeGroupMaxSize = 5000
	NodeGroupMinSize = 1
)

type DummyNodeGroup struct {
	nodes []apiv1.Node
}

var NodeGroup = &DummyNodeGroup{
	nodes: []apiv1.Node{},
}

func (dng *DummyNodeGroup) MaxSize() int {
	return NodeGroupMaxSize
}

func (dng *DummyNodeGroup) MinSize() int {
	return NodeGroupMinSize
}

func (dng *DummyNodeGroup) TargetSize() (int, error) {
	glog.V(1).Infof("Reporting TargetSize: %d", len(dng.nodes))
	return len(dng.nodes), nil
}

func (dng *DummyNodeGroup) IncreaseSize(delta int) error {
	glog.Infof("IncreaseSize invoked (delta=%d)", delta)
	if *dummyWebHook != "" {
		req, err := http.NewRequest("PUT", *dummyWebHook+fmt.Sprintf("/%d", delta), nil)
		if err != nil {
			glog.Errorf("PUT request cannot be forged for dummy webhook: %+v", err)
			return nil
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			glog.Errorf("error sending PUT request to dummy webhook: %+v", err)
			return nil
		}
		if resp.StatusCode > 299 || resp.StatusCode < 200 {
			glog.Errorf("dummy webhook returned %d", resp.StatusCode)
		}
	}
	return nil
}

func (dng *DummyNodeGroup) DeleteNodes(nodes []*apiv1.Node) error {
	glog.Infof("DeleteNodes invoked (%d nodes)", len(nodes))
	if *dummyWebHook != "" {
		for _, node := range nodes {
			req, err := http.NewRequest("DELETE", *dummyWebHook+fmt.Sprintf("/%s", node.Name), nil)
			if err != nil {
				glog.Errorf("DELETE request cannot be forged for dummy webhook: %+v", err)
				return nil
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				glog.Errorf("error sending DELETE request to dummy webhook: %+v", err)
				return nil
			}
			if resp.StatusCode > 299 || resp.StatusCode < 200 {
				glog.Errorf("dummy webhook returned %d", resp.StatusCode)
			}
		}
	}
	return nil
}

func (dng *DummyNodeGroup) DecreaseTargetSize(delta int) error {
	glog.Infof("DecreaseTargetSize invoked (delta=%d)", delta)
	return nil
}

func (dng *DummyNodeGroup) Id() string {
	return "dummy-node-group"
}

func (dng *DummyNodeGroup) Debug() string {
	return fmt.Sprintf("%+v", dng)
}

func (dng *DummyNodeGroup) Nodes() ([]string, error) {
	nodes := make([]string, len(dng.nodes))
	for i, node := range dng.nodes {
		nodes[i] = node.Name
	}
	glog.Infof("Nodes invoked (result=%+v)", nodes)
	return nodes, nil
}

func (dng *DummyNodeGroup) TemplateNodeInfo() (*schedulercache.NodeInfo, error) {
	return nil, nil
}

func (dng *DummyNodeGroup) Exist() bool {
	return true
}

func (dng *DummyNodeGroup) Create() error {
	return nil
}

func (dng *DummyNodeGroup) Delete() error {
	return nil
}

func (dng *DummyNodeGroup) Autoprovisioned() bool {
	return false
}
