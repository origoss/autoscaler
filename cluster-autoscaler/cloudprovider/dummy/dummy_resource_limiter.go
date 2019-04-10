package dummy

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

var DummyResourceLimiter = cloudprovider.NewResourceLimiter(map[string]int64{
	"dummy-limit": 0,
}, map[string]int64{
	"dummy-limit": 100,
})
