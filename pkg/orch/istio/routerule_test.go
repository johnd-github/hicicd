package istio

import (
	"testing"
	"github.com/hidevopsio/hicicd/pkg/orch"
	"github.com/stretchr/testify/assert"
	"github.com/hidevopsio/hiboot/pkg/log"
	routing "istio.io/api/routing/v1alpha1"
)

func TestCreate(t *testing.T) {
	rule := &Rule{
		Labels:          map[string]string{"label": "value"},
		Annotations:     map[string]string{"annotation": "value"},
		Name:            "hello-world",
		Namespace:       "demo-dev",
		Route: []*routing.DestinationWeight{
			{Weight: 80, Labels: map[string]string{"version": "v1"}},
			{Weight: 20, Labels: map[string]string{"version": "v3"}},
		},
	}
	config, err := NewClient(*orch.Kubeconfig)
	assert.Equal(t, nil, err)
	resourceVersion, err := rule.Create(config)
	assert.Equal(t, nil, err)
	log.Info("create rule:", resourceVersion)
}

func TestGet(t *testing.T) {
	config, err := NewClient(*orch.Kubeconfig)
	assert.Equal(t, nil, err)
	rule := Rule{
		Name:      "hello-world",
		Namespace: "demo-dev",
	}

	con, flag := rule.Get(config)
	log.Info("get rule :", con)
	resoureceVersion := con.ResourceVersion
	assert.Equal(t, true, flag)
	log.Info("Roter Rule get resourece version :",resoureceVersion)
}

func TestUpdate(t *testing.T) {
	rule := &Rule{
		Labels:          map[string]string{"label": "value"},
		Annotations:     map[string]string{"annotation": "value"},
		Name:            "hello-world",
		Namespace:       "demo-dev",
		ResourceVersion: "3897304",
		Route: []*routing.DestinationWeight{
			{Weight: 80, Labels: map[string]string{"version": "v1"}},
			{Weight: 20, Labels: map[string]string{"version": "v2"}},
		},
	}
	config, err := NewClient(*orch.Kubeconfig)
	assert.Equal(t, nil, err)
	resourceVersion, err := rule.Update(config)
	assert.Equal(t, nil, err)
	log.Info("update rule:", resourceVersion)
}

func TestDelete(t *testing.T) {
	config, err := NewClient(*orch.Kubeconfig)
	assert.Equal(t, nil, err)
	rule := Rule{
		Name:      "hello-world",
		Namespace: "demo-dev",
	}

	err = rule.Delete(config)
	assert.Equal(t, nil, err)
}