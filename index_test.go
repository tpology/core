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
	err := i.AddResource(&v1.Resource{
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
					File:       "path-1",
					Template:   "template-1",
				},
			},
		},
	})
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
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
	if len(res.Resource.Data.(map[string]interface{})) != 1 {
		t.Errorf("Expected 1 data, got %d", len(res.Resource.Data.(map[string]interface{})))
	}
	if res.Resource.Data.(map[string]interface{})["data"] != "value" {
		t.Errorf("Expected value, got %s", res.Resource.Data.(map[string]interface{})["data"])
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
	if res.Resource.Outputs[0].File != "path-1" {
		t.Errorf("Expected path-1, got %s", res.Resource.Outputs[0].File)
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
	err := i.AddResource(r)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 kind, got %d", len(i.resourceByKind))
	}
	err = i.RemoveResource(r)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.resourceByKind) != 0 {
		t.Errorf("Expected 0 kind, got %d", len(i.resourceByKind))
	}
}

// Test_Index_RemoveResource_Missing tests the RemoveResource function of the
// Index. It removes one Resource that does not exist and then checks that it
// was not removed.
func Test_Index_RemoveResource_Missing(t *testing.T) {
	i := NewIndex()
	r := &v1.Resource{
		APIVersion: "v1",
		Resource: v1.ResourceSpec{
			Name: "resource-1",
			Kind: "test",
		},
	}
	err := i.RemoveResource(r)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "resource resource-1 of kind test does not exist" {
		t.Errorf("Expected error 'resource resource-1 of kind test does not exist', got %s", err.Error())
	}
}

// Test_Index_AddTemplate tests the AddTemplate function of the Index. It
// adds one Template and then checks that it was added.
func Test_Index_AddTemplate(t *testing.T) {
	i := NewIndex()
	err := i.AddTemplate(&v1.Template{
		APIVersion: "v1",
		Template: v1.TemplateSpec{
			Name:    "template-1",
			Content: "test",
		},
	})
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
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
	err := i.AddTemplate(tmpl)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.template) != 1 {
		t.Errorf("Expected 1 template, got %d", len(i.template))
	}
	err = i.RemoveTemplate(tmpl)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.template) != 0 {
		t.Errorf("Expected 0 template, got %d", len(i.template))
	}
}

// Test_Index_RemoveTemplate_Missing tests the RemoveTemplate function of the
// Index. It removes one Template that does not exist and then checks that it
// was not removed.
func Test_Index_RemoveTemplate_Missing(t *testing.T) {
	i := NewIndex()
	tmpl := &v1.Template{
		APIVersion: "v1",
		Template: v1.TemplateSpec{
			Name:    "template-1",
			Content: "test",
		},
	}
	err := i.RemoveTemplate(tmpl)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "template template-1 does not exist" {
		t.Errorf("Expected error 'template template-1 does not exist', got %s", err.Error())
	}
}

// Test_Index_AddRepository tests the AddRepository function of the Index. It
// adds one Repository and then checks that it was added.
func Test_Index_AddRepository(t *testing.T) {
	i := NewIndex()
	err := i.AddRepository(&v1.Repository{
		APIVersion: "v1",
		Repository: v1.RepositorySpec{
			Name:       "repo-1",
			Repository: "test-repo-1",
			Branch:     "test-branch",
		},
	})
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.repository) != 1 {
		t.Errorf("Expected 1 repository, got %d", len(i.repository))
	}
	repo := i.repository["repo-1"]
	if repo.Repository.Name != "repo-1" {
		t.Errorf("Expected repo-1, got %s", repo.Repository.Name)
	}
	if repo.Repository.Repository != "test-repo-1" {
		t.Errorf("Expected test-repo-1, got %s", repo.Repository.Repository)
	}
	if repo.Repository.Branch != "test-branch" {
		t.Errorf("Expected test-branch, got %s", repo.Repository.Branch)
	}
	if len(repo.Repository.Labels) != 0 {
		t.Errorf("Expected 0 labels, got %d", len(repo.Repository.Labels))
	}
	if len(repo.Repository.Annotations) != 0 {
		t.Errorf("Expected 0 annotations, got %d", len(repo.Repository.Annotations))
	}
}

// Test_Index_RemoveRepository tests the RemoveRepository function of the
// Index. It removes one Repository and then checks that it was removed.
func Test_Index_RemoveRepository(t *testing.T) {
	i := NewIndex()
	repo := &v1.Repository{
		APIVersion: "v1",
		Repository: v1.RepositorySpec{
			Name:       "repo-1",
			Repository: "test-repo-1",
			Branch:     "test-branch",
		},
	}
	err := i.AddRepository(repo)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.repository) != 1 {
		t.Errorf("Expected 1 repository, got %d", len(i.repository))
	}
	err = i.RemoveRepository(repo)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(i.repository) != 0 {
		t.Errorf("Expected 0 repository, got %d", len(i.repository))
	}
}

// Test_Index_RemoveRepository_Missing tests the RemoveRepository function of
// the Index. It removes one Repository that does not exist and then checks
// that it was not removed.
func Test_Index_RemoveRepository_Missing(t *testing.T) {
	i := NewIndex()
	repo := &v1.Repository{
		APIVersion: "v1",
		Repository: v1.RepositorySpec{
			Name:       "repo-1",
			Repository: "test-repo-1",
			Branch:     "test-branch",
		},
	}
	err := i.RemoveRepository(repo)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "repository repo-1 does not exist" {
		t.Errorf("Expected error 'repository repo-1 does not exist', got %s", err.Error())
	}
}

// Test_Index_Load tests the Load function of the Index
func Test_Index_Load_Basic_Resource(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/000-basic-resource", nil)
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
	if res.Resource.Data != nil {
		t.Errorf("Expected nil data, got %v", res.Resource.Data)
	}
	if len(res.Resource.Outputs) != 0 {
		t.Errorf("Expected 0 outputs, got %d", len(res.Resource.Outputs))
	}
}

// Test_Index_Load_Basic_Template tests the Load function of the Index
func Test_Index_Load_Basic_Template(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/001-basic-template", nil)
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

// Test_Index_Load_Basic_Repository tests the Load function of the Index
func Test_Index_Load_Basic_Repository(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/016-basic-repository", nil)
	if len(i.resourceByKind) != 0 {
		t.Errorf("Expected 0 kinds, got %d", len(i.resourceByKind))
	}
	if len(i.template) != 0 {
		t.Errorf("Expected 0 templates, got %d", len(i.template))
	}
	if len(i.repository) != 1 {
		t.Errorf("Expected 1 repository, got %d", len(i.repository))
	}
	repo := i.repository["repo-1"]
	if repo.Repository.Name != "repo-1" {
		t.Errorf("Expected repo-1, got %s", repo.Repository.Name)
	}
	if repo.Repository.Repository != "test-repo-1" {
		t.Errorf("Expected test-repo-1, got %s", repo.Repository.Repository)
	}
	if repo.Repository.Branch != "test-branch" {
		t.Errorf("Expected test-branch, got %s", repo.Repository.Branch)
	}
	if len(repo.Repository.Labels) != 0 {
		t.Errorf("Expected 0 labels, got %d", len(repo.Repository.Labels))
	}
	if len(repo.Repository.Annotations) != 0 {
		t.Errorf("Expected 0 annotations, got %d", len(repo.Repository.Annotations))
	}
}

// Test_Index_Load_TwoResources tests the Load function of the Index. It
// expects to find one Resource and one Template.
func Test_Index_Load_TwoResources(t *testing.T) {
	i := NewIndex()
	i.Load("testdata/002-two-resources", nil)
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

// mustChmod is a helper function that chmods a file and panics on error.
func mustChmod(t *testing.T, path string, mode os.FileMode) {
	if err := os.Chmod(path, mode); err != nil {
		t.Fatalf("Failed to chmod %s to %s: %v", path, mode, err)
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
	defer mustChmod(t, "testdata/003-load-unreadable-file/resource-1.yaml", 0664)
	i := NewIndex()
	errs := i.Load("testdata/003-load-unreadable-file", nil)
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
	defer mustChmod(t, "testdata/004-load-unreadable-dir", 0775)
	i := NewIndex()
	errs := i.Load("testdata/004-load-unreadable-dir", nil)
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
	errs := i.Load("testdata/005-load-corrupted-document", nil)
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
	errs := i.Load("testdata/006-missing-apiversion", nil)
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
	errs := i.Load("testdata/007-invalid-apiversion", nil)
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
	errs := i.Load("testdata/008-invalid-resource", nil)
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/008-invalid-resource/resource-1.yaml: no resource or template" {
		t.Errorf("Expected testdata/008-invalid-resource/resource-1.yaml: no resource or template, got %s", errs[0].Error())
	}
}

// Test_Index_Load_MultipleErrors tests the Load function of the Index. It
// expects to receive multiple errors.
func Test_Index_Load_MultipleErrors(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/017-multiple-errors", nil)
	if len(errs) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/017-multiple-errors/resource-1.yaml: no apiVersion" {
		t.Errorf("Expected testdata/017-multiple-errors/resource-1.yaml: no apiVersion, got %s", errs[0].Error())
	}
	if errs[1].Error() != "testdata/017-multiple-errors/resource-2.yaml: invalid apiVersion" {
		t.Errorf("Expected testdata/017-multiple-errors/resource-2.yaml: invalid apiVersion, got %s", errs[1].Error())
	}
}

// Test_Index_LoadResourceWithOutput tests the Load function of the Index. It
// expects to load 1 Resource with 1 Output.
func Test_Index_LoadResourceWithOutput(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/023-resource-with-output", nil)
	if len(errs) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(errs))
	}
	if len(i.resourceByKind) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind))
	}
	if len(i.resourceByKind["test"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test"]))
	}
	resource := i.resourceByKind["test"]["resource-1"]
	if resource.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", resource.Resource.Name)
	}
	if len(resource.Resource.Outputs) != 1 {
		t.Errorf("Expected 1 output, got %d", len(resource.Resource.Outputs))
	}
	// Validate name
	if resource.Resource.Outputs[0].Name != "output-1" {
		t.Errorf("Expected output-1, got %s", resource.Resource.Outputs[0].Name)
	}
	// Validate path
	if resource.Resource.Outputs[0].File != "path" {
		t.Errorf("Expected path, got %s", resource.Resource.Outputs[0].File)
	}
	// Validate repository
	if resource.Resource.Outputs[0].Repository != "repository" {
		t.Errorf("Expected repository, got %s", resource.Resource.Outputs[0].Repository)
	}
	// Validate template
	if resource.Resource.Outputs[0].Template != "template" {
		t.Errorf("Expected template, got %s", resource.Resource.Outputs[0].Template)
	}
	// Validate postProcessor
	if resource.Resource.Outputs[0].PostProcessor != "postProcessor" {
		t.Errorf("Expected postProcessor, got %s", resource.Resource.Outputs[0].PostProcessor)
	}
}

// Test_Index_Load_TwoResourceSameName tests the Load function of the Index. It
// expects to get an error trying to load 2 resources of the same kind and name.
func Test_Index_Load_TwoResourceSameName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/024-two-resources-same-name", nil)
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/024-two-resources-same-name/resource-2.yaml: resource resource-1 of kind test already exists" {
		t.Errorf("Expected testdata/024-two-resources-same-name/resource-2.yaml: resource resource-1 of kind test already exists, got %s", errs[0].Error())
	}
}

// Test_Index_Load_TwoResourceSameNameDifferentKind tests the Load function of
// the Index. It expects to successfully load two resources of the same name
// but different kinds.
func Test_Index_Load_TwoResourceSameNameDifferentKind(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/025-two-resources-same-name-different-kind", nil)
	if len(errs) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(errs))
	}
	if len(i.resourceByKind) != 2 {
		t.Errorf("Expected 2 resources, got %d", len(i.resourceByKind))
	}
	if len(i.resourceByKind["test"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test"]))
	}
	if len(i.resourceByKind["test2"]) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(i.resourceByKind["test2"]))
	}
	resource := i.resourceByKind["test"]["resource-1"]
	if resource.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", resource.Resource.Name)
	}
	resource = i.resourceByKind["test2"]["resource-1"]
	if resource.Resource.Name != "resource-1" {
		t.Errorf("Expected resource-1, got %s", resource.Resource.Name)
	}
}

// Test_Index_Load_TwoTemplateSameName tests the Load function of the Index. It
// expects to get an error trying to load 2 templates of the same name.
func Test_Index_Load_TwoTemplateSameName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/026-two-templates-same-name", nil)
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/026-two-templates-same-name/template-2.yaml: template template-1 already exists" {
		t.Errorf("Expected testdata/026-two-templates-same-name/template-2.yaml: template template-1 already exists, got %s", errs[0].Error())
	}
}

// Test_Index_Load_TwoRepositorySameName tests the Load function of the Index.
// It expects to get an error trying to load 2 repositories of the same name.
func Test_Index_Load_TwoRepositorySameName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/027-two-repositories-same-name", nil)
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/027-two-repositories-same-name/repository-2.yaml: repository repo-1 already exists" {
		t.Errorf("Expected testdata/027-two-repositories-same-name/repository-2.yaml: repository repo-1 already exists, got %s", errs[0].Error())
	}
}

// Test_Index_LoadNonexistentDir tests the Load function of the Index. It
// expects to get an error trying to load a directory that doesn't exist.
func Test_Index_LoadNonexistentDir(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/nonexistent-dir", nil)
	if len(errs) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "lstat testdata/nonexistent-dir: no such file or directory" {
		t.Errorf("Expected lstat testdata/nonexistent-dir: no such file or directory, got %s", errs[0].Error())
	}
}
