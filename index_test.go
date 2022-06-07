package core

import "testing"

// Test_Index_Load tests the Load function of the Index
func Test_Index_Load_Basic_Resource(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/000-basic-resource")
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 kind, got %d", len(i.resourceByKind))
	}
	if len(i.resourceByKind["test"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test"]))
	}
	if len(i.template) != 0 {
		t.Errorf("Expected 0 templates, got %d", len(i.template))
	}
	res := i.resourceByKind["test"]["resource-1"]
	if res.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", res.Resource.Name)
	}
	if res.Resource.Kind != "test" {
		t.Errorf("Expected test, got %s", res.Resource.Kind)
	}
	if len(res.Resource.Labels) != 0 {
		t.Errorf("Expected 0 labels, got %d", len(res.Resource.Labels))
	}
	if len(res.Resource.Annotations) != 0 {
		t.Errorf("Expected 0 annotations, got %d", len(res.Resource.Annotations))
	}
	if len(res.Resource.Data) != 0 {
		t.Errorf("Expected 0 data, got %d", len(res.Resource.Data))
	}
	if len(res.Resource.Outputs) != 0 {
		t.Errorf("Expected 0 outputs, got %d", len(res.Resource.Outputs))
	}
}
