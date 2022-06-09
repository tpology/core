package core

import (
	"testing"

	v1 "github.com/tpology/core/api/v1"
)

// Test_Validate_MissingResourceKind tests that a resource without a kind is not valid
func Test_Validate_MissingResourceKind(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/009-missing-resource-kind", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/009-missing-resource-kind/resource-1.yaml: resource kind is required" {
		t.Errorf("expected 'testdata/009-missing-resource-kind/resource-1.yaml: resource kind is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingResourceName tests that a resource without a name is not valid
func Test_Validate_MissingResourceName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/010-missing-resource-name", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/010-missing-resource-name/resource-1.yaml: resource name is required" {
		t.Errorf("expected 'testdata/010-missing-resource-name/resource-1.yaml: resource name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingTemplateName tests that a template without a name is not valid
func Test_Validate_MissingTemplateName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/011-missing-template-name", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/011-missing-template-name/template-1.yaml: template name is required" {
		t.Errorf("expected 'testdata/011-missing-template-name/template-1.yaml: template name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidResourceField tests that an invalid resource field is not valid
func Test_Validate_InvalidResourceField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/012-invalid-resource-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/012-invalid-resource-field/resource-1.yaml: invalid resource field `invalid`" {
		t.Errorf("expected 'testdata/012-invalid-resource-field/resource-1.yaml: invalid resource field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidResourceSpecField tests that an invalid resource spec field is not valid
func Test_Validate_InvalidResourceSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/013-invalid-resource-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/013-invalid-resource-spec-field/resource-1.yaml: invalid resource spec field `invalid`" {
		t.Errorf("expected 'testdata/013-invalid-resource-spec-field/resource-1.yaml: invalid resource spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateField tests that an invalid template field is not valid
func Test_Validate_InvalidTemplateField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/014-invalid-template-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/014-invalid-template-field/template-1.yaml: invalid template field `invalid`" {
		t.Errorf("expected 'testdata/014-invalid-template-field/template-1.yaml: invalid template field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidTemplateSpecField tests that an invalid template spec field is not valid
func Test_Validate_InvalidTemplateSpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/015-invalid-template-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/015-invalid-template-spec-field/template-1.yaml: invalid template spec field `invalid`" {
		t.Errorf("expected 'testdata/015-invalid-template-spec-field/template-1.yaml: invalid template spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidRepositoryField tests that an invalid repository field is not valid
func Test_Validate_InvalidRepositoryField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/021-invalid-repository-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/021-invalid-repository-field/repository-1.yaml: invalid repository field `invalid`" {
		t.Errorf("expected 'testdata/021-invalid-repository-field/repository-1.yaml: invalid repository field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidRepositorySpecField tests that an invalid repository spec field is not valid
func Test_Validate_InvalidRepositorySpecField(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/022-invalid-repository-spec-field", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/022-invalid-repository-spec-field/repository-1.yaml: invalid repository spec field `invalid`" {
		t.Errorf("expected 'testdata/022-invalid-repository-spec-field/repository-1.yaml: invalid repository spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepositoryName tests that a repository without a name is not valid
func Test_Validate_MissingRepositoryName(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/018-missing-repository-name", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/018-missing-repository-name/repository-1.yaml: repository name is required" {
		t.Errorf("expected 'testdata/018-missing-repository-name/repository-1.yaml: repository name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepository tests that a repository without a repo is not valid
func Test_Validate_MissingRepository(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/019-missing-repository", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/019-missing-repository/repository-1.yaml: repository is required" {
		t.Errorf("expected 'testdata/019-missing-repository/repository-1.yaml: repository is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_MissingRepositoryBranch tests that a repository without a branch is not valid
func Test_Validate_MissingRepositoryBranch(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/020-missing-repository-branch", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/020-missing-repository-branch/repository-1.yaml: repository branch is required" {
		t.Errorf("expected 'testdata/020-missing-repository-branch/repository-1.yaml: repository branch is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidOutputSpecFields tests that an invalid output spec field is not valid
func Test_Validate_InvalidOutputSpecFields(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/028-invalid-outputspec-fields", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/028-invalid-outputspec-fields/resource-1.yaml: invalid output spec field `invalid`" {
		t.Errorf("expected 'testdata/028-invalid-outputspec-fields/resource-1.yaml: invalid output spec field `invalid`', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidOutputSpec_Name tests that an invalid output spec name is not valid
func Test_Validate_InvalidOutputSpec_Name(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/029-invalid-outputspec-name", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/029-invalid-outputspec-name/resource-1.yaml: output name is required" {
		t.Errorf("expected 'testdata/029-invalid-outputspec-name/resource-1.yaml: output name is required', got '%s'", errs[0].Error())
	}
}

// Test_Valid_InvalidOutputSpec_Template tests that an invalid output spec template is not valid
func Test_Validate_InvalidOutputSpec_Template(t *testing.T) {
	i := NewIndex()
	errs := i.Load("testdata/030-invalid-outputspec-template", nil)
	if len(errs) != 1 {
		t.Errorf("expected 1 error, got %d", len(errs))
	}
	if errs[0].Error() != "testdata/030-invalid-outputspec-template/resource-1.yaml: output template is required" {
		t.Errorf("expected 'testdata/030-invalid-outputspec-template/resource-1.yaml: output template is required', got '%s'", errs[0].Error())
	}
}

// Test_ValidateLoadPolicy_Allow tests that a valid policy is allowed
func Test_ValidateLoadPolicy_Allow(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "name",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
}

// Test_ValidateLoadPolicy_Deny tests that a valid policy is allowed
func Test_ValidateLoadPolicy_Deny(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "name",
		Effect: v1.Deny,
		Paths: []string{
			"/",
		},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
}

// Test_ValidateLoadPolicy_InvalidName tests that an invalid name is not valid
func Test_ValidateLoadPolicy_InvalidName(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy name is required" {
		t.Errorf("expected 'load policy name is required', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_InvalidEffect tests that an invalid effect is not valid
func Test_ValidateLoadPolicy_InvalidEffect(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: "invalid",
		Paths: []string{
			"/",
		},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy effect must be either `Allow` or `Deny`" {
		t.Errorf("expected 'load policy effect must be either `Allow` or `Deny`', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_EmptyEffect tests that an empty effect is not valid
func Test_ValidateLoadPolicy_EmptyEffect(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: "",
		Paths: []string{
			"/",
		},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy effect must be either `Allow` or `Deny`" {
		t.Errorf("expected 'load policy effect must be either `Allow` or `Deny`', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_NoPaths tests that an empty paths is not valid
func Test_ValidateLoadPolicy_NoPaths(t *testing.T) {
	p := v1.LoadPolicy{
		Name:       "test",
		Effect:     v1.Allow,
		Paths:      []string{},
		Resource:   &v1.ResourceSpec{},
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy requires at least one path" {
		t.Errorf("expected 'load policy requires at least one path', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_NoResourceTemplateOrRepository tests that an empty resource template or repository is not valid
func Test_ValidateLoadPolicy_NoResourceTemplateOrRepository(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource:   nil,
		Template:   nil,
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy must have one of template, repository, or resource" {
		t.Errorf("expected 'load policy must have one of template, repository, or resource', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_ResourceAndTemplate_Invalid tests that an invalid resource and template is not valid
func Test_ValidateLoadPolicy_ResourceAndTemplate_Invalid(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource: &v1.ResourceSpec{
			Name: "test",
		},
		Template: &v1.TemplateSpec{
			Name: "test",
		},
		Repository: nil,
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy must have exclusively one template, repository, or resource" {
		t.Errorf("expected 'load policy must have exclusively one template, repository, or resource', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_ResourceAndRepository_Invalid tests that an invalid resource and repository is not valid
func Test_ValidateLoadPolicy_ResourceAndRepository_Invalid(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource: &v1.ResourceSpec{
			Name: "test",
		},
		Template: nil,
		Repository: &v1.RepositorySpec{
			Name: "test",
		},
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy must have exclusively one template, repository, or resource" {
		t.Errorf("expected 'load policy must have exclusively one template, repository, or resource', got '%s'", err.Error())
	}
}

// Test_ValidateLoadPolicy_RepositoryAndTemplate_Invalid tests that an invalid repository and template is not valid
func Test_ValidateLoadPolicy_RepositoryAndTemplate_Invalid(t *testing.T) {
	p := v1.LoadPolicy{
		Name:   "test",
		Effect: v1.Allow,
		Paths: []string{
			"/",
		},
		Resource: nil,
		Template: &v1.TemplateSpec{
			Name: "test",
		},
		Repository: &v1.RepositorySpec{
			Name: "test",
		},
	}
	err := validateLoadPolicy(p)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "load policy must have exclusively one template, repository, or resource" {
		t.Errorf("expected 'load policy must have exclusively one template, repository, or resource', got '%s'", err.Error())
	}
}
