/*
Copyright 2020-2023 Oracle and/or its affiliates.
*/

package nodepools

import (
	"context"
	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	ocicommon "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/common"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestDeletePastMinSize(t *testing.T) {
	client := fake.NewSimpleClientset()

	nodeNames := []string{
		"nodeA",
	}
	manager := &mockManager{
		err:        nil,
		timeOutErr: apierrors.NewTimeoutError("timeout error", 5),
		nodePoolSize: 1,
	}
	np := &nodePool{
		kubeClient: client,
		manager:    manager,
		minSize:    1,
		maxSize:    10,
		id:         "abc",
	}
	manager.nodePool = np

	var nodesToDelete []*apiv1.Node
	for _, name := range nodeNames {
		node, err := client.CoreV1().Nodes().Create(context.Background(), &apiv1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("unexpected create error: %+v", err)
		}
		nodesToDelete = append(nodesToDelete, node)
	}
	err := np.DeleteNodes(nodesToDelete)
	if err == nil {
		t.Fatalf("expected to have an error because node pool is at the min size")
	}
}

func TestDeleteNodeWithoutInstanceIDFallsBackToDecreaseTargetSize(t *testing.T) {
	client := fake.NewSimpleClientset()

	manager := &mockManager{
		nodePoolSize: 1,
		nodes: []cloudprovider.Instance{
			{
				Id: "",
				Status: &cloudprovider.InstanceStatus{
					State: cloudprovider.InstanceCreating,
					ErrorInfo: &cloudprovider.InstanceErrorInfo{
						ErrorClass:   cloudprovider.OutOfResourcesErrorClass,
						ErrorCode:    "InternalError",
						ErrorMessage: "Out of host capacity",
					},
				},
			},
		},
	}

	np := &nodePool{
		kubeClient: client,
		manager:    manager,
		minSize:    0,
		maxSize:    10,
		id:         "abc",
	}
	manager.nodePool = np

	node := &apiv1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "badNode",
		},
	}

	err := np.ForceDeleteNodes([]*apiv1.Node{node})
	if err != nil {
		t.Fatalf("expected fallback cleanup to succeed, got: %v", err)
	}
}

type mockManager struct {
	called    []string
	nodePools []NodePool
	nodes     []cloudprovider.Instance

	// used for GetNodePoolForInstance
	nodePool NodePool
	NodePoolManager
	err        error
	timeOutErr error

	nodePoolSize int
}

func (m mockManager) Refresh() error {
	m.called = append(m.called, "refresh")
	return nil
}

func (m mockManager) Cleanup() error {
	m.called = append(m.called, "cleanup")
	return nil
}

func (m mockManager) GetNodePools() []NodePool {
	m.called = append(m.called, "get-node-pools")
	return m.nodePools
}

func (m mockManager) GetNodePoolNodes(np NodePool) ([]cloudprovider.Instance, error) {
	m.called = append(m.called, "get-node-pool-nodes")
	return m.nodes, nil
}

func (m mockManager) GetNodePoolForInstance(instance ocicommon.OciRef) (NodePool, error) {
	m.called = append(m.called, "get-node-pool-for-instance")
	return m.nodePool, m.err
}

func (m mockManager) GetNodePoolTemplateNode(np NodePool) (*apiv1.Node, error) {
	m.called = append(m.called, "get-node-pool-template-node")
	panic("implement me")
}

func (m mockManager) GetResourceLimiter() (*cloudprovider.ResourceLimiter, error) {
	m.called = append(m.called, "get-resource-limiter")
	panic("implement me")
}

func (m mockManager) GetNodePoolSize(np NodePool) (int, error) {
	m.called = append(m.called, "get-node-pool-size")
	if m.nodePoolSize != 0 {
		return m.nodePoolSize, nil
	}
	return np.MinSize() + 1, nil
}

func (m mockManager) SetNodePoolSize(np NodePool, size int) error {
	m.called = append(m.called, "set-node-pool-size")
	return nil
}

func (m mockManager) DeleteInstances(np NodePool, instances []ocicommon.OciRef) error {
	m.called = append(m.called, "delete-instances")
	return m.timeOutErr
}

func (m mockManager) InvalidateAndRefreshCache() error {
	m.called = append(m.called, "invalidate-and-refresh-cache")
	return nil
}

func (m mockManager) GetExistingNodePoolSizeViaCompute(np NodePool) (int, error) {
	m.called = append(m.called, "get-existing-node-pool-size-via-compute")
	return 0, nil
}
