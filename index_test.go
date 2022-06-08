package core

import (
	"os"
	"testing"

	v1 "github.com/tpology/core/api/v1"
)

// Test_Index_AddResource tests the AddResource function of the Index. It
// adds one Resource and then checks that it was added.
func Test_Index_AddResource(t *testing.T) {
	i := NewIndex()
	i.AddResource(&v1.Resource{
		APIVersion: "v1",
		Resource: v1.ResourceSpec{
			Name:        "resource-1",
			Kind:        "test",
			Labels:      map[string]string{"label": "value"},
			Annotations: map[string]string{"annotation": "value"},
			Data:        map[string]interface{}{"data": "value"},
			Outputs: []v1.OutputSpec{
				{
					Name:       "output-1",
					Repository: "repo-1",
					Path:       "path-1",
					Template:   "template-1",
				},
			},
		},
	})
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 kind, got %d", len(i.resourceByKind))
	}
	if len(i.resourceByKind["test"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test"]))
	}
	res := i.resourceByKind["test"]["resource-1"]
	if res.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", res.Resource.Name)
	}
	if res.Resource.Kind != "test" {
		t.Errorf("Expected test, got %s", res.Resource.Kind)
	}
	if len(res.Resource.Labels) != 1 {
		t.Errorf("Expected 1 label, got %d", len(res.Resource.Labels))
	}
	if res.Resource.Labels["label"] != "value" {
		t.Errorf("Expected value, got %s", res.Resource.Labels["label"])
	}
	if len(res.Resource.Annotations) != 1 {
		t.Errorf("Expected 1 annotation, got %d", len(res.Resource.Annotations))
	}
	if res.Resource.Annotations["annotation"] != "value" {
		t.Errorf("Expected value, got %s", res.Resource.Annotations["annotation"])
	}
	if len(res.Resource.Data) != 1 {
		t.Errorf("Expected 1 data, got %d", len(res.Resource.Data))
	}
	if res.Resource.Data["data"] != "value" {
		t.Errorf("Expected value, got %s", res.Resource.Data["data"])
	}
	if len(res.Resource.Outputs) != 1 {
		t.Errorf("Expected 1 output, got %d", len(res.Resource.Outputs))
	}
	if res.Resource.Outputs[0].Name != "output-1" {
		t.Errorf("Expected output-1, got %s", res.Resource.Outputs[0].Name)
	}
	if res.Resource.Outputs[0].Repository != "repo-1" {
		t.Errorf("Expected repo-1, got %s", res.Resource.Outputs[0].Repository)
	}
	if res.Resource.Outputs[0].Path != "path-1" {
		t.Errorf("Expected path-1, got %s", res.Resource.Outputs[0].Path)
	}
	if res.Resource.Outputs[0].Template != "template-1" {
		t.Errorf("Expected template-1, got %s", res.Resource.Outputs[0].Template)
	}
}

// Test_Index_RemoveResource tests the RemoveResource function of the Index. It
// removes one Resource and then checks that it was removed.
func Test_Index_RemoveResource(t *testing.T) {
	i := NewIndex()
	r := &v1.Resource{
		APIVersion: "v1",
		Resource: v1.ResourceSpec{
			Name: "resource-1",
			Kind: "test",
		},
	}
	i.AddResource(r)
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 kind, got %d", len(i.resourceByKind))
	}
	i.RemoveResource(r)
	if len(i.resourceByKind) != 0 {
		t.Errorf("Expected 0 kind, got %d", len(i.resourceByKind))
	}
}

// Test_Index_AddTemplate tests the AddTemplate function of the Index. It
// adds one Template and then checks that it was added.
func Test_Index_AddTemplate(t *testing.T) {
	i := NewIndex()
	i.AddTemplate(&v1.Template{
		APIVersion: "v1",
		Template: v1.TemplateSpec{
			Name:    "template-1",
			Content: "test",
		},
	})
	if len(i.template) != 1 {
		t.Errorf("Expected 1 template, got %d", len(i.template))
	}
	tpl := i.template["template-1"]
	if tpl.Template.Name != "template-1" {
		t.Errorf("Expected template-1, got %s", tpl.Template.Name)
	}
	if tpl.Template.Content != "test" {
		t.Errorf("Expected test, got %s", tpl.Template.Content)
	}
}

// Test_Index_RemoveTemplate tests the RemoveTemplate function of the Index. It
// removes one Template and then checks that it was removed.
func Test_Index_RemoveTemplate(t *testing.T) {
	i := NewIndex()
	tmpl := &v1.Template{
		APIVersion: "v1",
		Template: v1.TemplateSpec{
			Name:    "template-1",
			Content: "test",
		},
	}
	i.AddTemplate(tmpl)
	if len(i.template) != 1 {
		t.Errorf("Expected 1 template, got %d", len(i.template))
	}
	i.RemoveTemplate(tmpl)
	if len(i.template) != 0 {
		t.Errorf("Expected 0 template, got %d", len(i.template))
	}
}

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

// Test_Index_Load_Basic_Template tests the Load function of the Index
func Test_Index_Load_Basic_Template(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/001-basic-template")
	if len(i.resourceByKind) != 0 {
		t.Errorf("Expected 0 kinds, got %d", len(i.resourceByKind))
	}
	if len(i.template) != 1 {
		t.Errorf("Expected 1 template, got %d", len(i.template))
	}
	tpl := i.template["template-1"]
	if tpl.Template.Name != "template-1" {
		t.Errorf("Expected template-1, got %s", tpl.Template.Name)
	}
	if tpl.Template.Content != "test" {
		t.Errorf("Expected test, got %s", tpl.Template.Content)
	}
}

// Test_Index_Load_TwoResources tests the Load function of the Index. It
// expects to find one Resource and one Template.
func Test_Index_Load_TwoResources(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/002-two-resources")
	// Expect 1 resource of kind test and name resource-1
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 kind, got %d", len(i.resourceByKind))
	}
	if len(i.resourceByKind["test"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test"]))
	}
	res := i.resourceByKind["test"]["resource-1"]
	if res.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", res.Resource.Name)
	}
	if res.Resource.Kind != "test" {
		t.Errorf("Expected test, got %s", res.Resource.Kind)
	}
	// Expect 1 template named template-1 with content "test"
	if len(i.template) != 1 {
		t.Errorf("Expected 1 template, got %d", len(i.template))
	}
	tpl := i.template["template-1"]
	if tpl.Template.Name != "template-1" {
		t.Errorf("Expected template-1, got %s", tpl.Template.Name)
	}
	if tpl.Template.Content != "test" {
		t.Errorf("Expected test, got %s", tpl.Template.Content)
	}
}

// Test_Index_UnreadableFile tests the Load function of the Index. It
// expects to receive an error due to an unreadable file.
func Test_Index_UnreadableFile(t *testing.T) {
	// Make the 003-load-unreadable-file/resource-1.yaml file unreadable
	err := os.Chmod("testdata/003-load-unreadable-file/resource-1.yaml", 0000)
	if err != nil {
		t.Errorf("Failed to chmod file: %s", err)
	}
	defer os.Chmod("testdata/003-load-unreadable-file/resource-1.yaml", 0664)
	i := NewIndex()
	errs := i.Load("testdata/003-load-unreadable-file")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "open testdata/003-load-unreadable-file/resource-1.yaml: permission denied" {
		t.Errorf("Expected open testdata/003-load-unreadable-file/resource-1.yaml: permission denied, got %s", errs[0].Error())
	}
}

// Test_Index_UnreadableDir tests the Load function of the Index. It
// expects to receive an error due to an unreadable directory.
func Test_Index_UnreadableDir(t *testing.T) {
	// Make the 003-load-unreadable-dir/ directory unreadable
	err := os.Chmod("testdata/004-load-unreadable-dir", 0000)
	if err != nil {
		t.Errorf("Failed to chmod directory: %s", err)
	}
	defer os.Chmod("testdata/004-load-unreadable-dir", 0775)
	i := NewIndex()
	errs := i.Load("testdata/004-load-unreadable-dir")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "open testdata/004-load-unreadable-dir: permission denied" {
		t.Errorf("Expected permission denied, got %s", errs[0].Error())
	}
}

// Test_Index_Load_CorruptedDocument tests the Load function of the Index. It
// expects to receive an error due to a corrupted document.
func Test_Index_Load_CorruptedDocument(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/005-load-corrupted-document")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/005-load-corrupted-document/resource-1.yaml: yaml: line 4: mapping values are not allowed in this context" {
		t.Errorf("Expected testdata/005-load-corrupted-document/resource-1.yaml: yaml: line 4: mapping values are not allowed in this context, got %s", errs[0].Error())
	}
}

// Test_Index_MissingAPIVersion tests the Load function of the Index. It
// expects to receive an error due to a missing API version.
func Test_Index_MissingAPIVersion(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/006-missing-apiversion")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/006-missing-apiversion/resource-1.yaml: no apiVersion" {
		t.Errorf("Expected testdata/006-missing-apiversion/resource-1.yaml: no apiVersion, got %s", errs[0].Error())
	}
}

// Test_Index_InvalidAPIVersion tests the Load function of the Index. It
// expects to receive an error due to an invalid API version.
func Test_Index_InvalidAPIVersion(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/007-invalid-apiversion")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/007-invalid-apiversion/resource-1.yaml: invalid apiVersion" {
		t.Errorf("Expected testdata/007-invalid-apiversion/resource-1.yaml: invalid apiVersion, got %s", errs[0].Error())
	}
}

// Test_Index_Load_InvalidResource tests the Load function of the Index. It
// expects to receive an error due to an invalid resource.
func Test_Index_Load_InvalidResource(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/008-invalid-resource")
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/008-invalid-resource/resource-1.yaml: no resource or template" {
		t.Errorf("Expected testdata/008-invalid-resource/resource-1.yaml: no resource or template, got %s", errs[0].Error())
	}
}
